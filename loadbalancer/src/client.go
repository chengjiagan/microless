package loadbalancer

import (
	"context"
	"log"
	"math/rand"
	"microless/loadbalancer/internal/utils"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"k8s.io/client-go/kubernetes"
)

type ClientLB struct {
	enable          bool
	vm              string
	serverless      string
	port            string
	degradeInterval int

	kubeClient *kubernetes.Clientset

	mu sync.RWMutex
	// mu protects the following fields
	vmConn         []*grpc.ClientConn
	serverlessConn []*grpc.ClientConn
	// mu read lock protects the following fields
	degradeConn    []uint32 // time left being degraded, >0: degraded, 0: normal
	numDegradConn  int32
	vmNext         uint32
	serverlessNext uint32
}

func splitAddr(addr string) (string, string) {
	splits := strings.Split(addr, ":")
	return splits[0], splits[1]
}

func NewClientLB(addr string) *ClientLB {
	config := utils.GetClientConfig()

	service, port := splitAddr(addr)
	vmServiceName := service + config.VmPostfix
	serverlessServiceName := service + config.ServerlessPostfix

	kubeClient, err := utils.NewKubeClient()
	if err != nil {
		log.Fatal(err)
	}

	lb := &ClientLB{
		enable:          config.Enable,
		port:            port,
		vm:              vmServiceName,
		serverless:      serverlessServiceName,
		degradeInterval: config.DegradeInterval * 1000, // seconds to milliseconds
		kubeClient:      kubeClient,
		vmNext:          rand.Uint32(),
		serverlessNext:  rand.Uint32(),
	}
	go lb.watchService()
	go lb.updateDegrade()

	return lb
}

func closeConn(conn []*grpc.ClientConn) {
	for _, c := range conn {
		err := c.Close()
		if err != nil {
			log.Printf("Failed to close conn: %v", err)
		}
	}
}

func (lb *ClientLB) updateVmConn(conn []*grpc.ClientConn) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	go closeConn(lb.vmConn)
	lb.vmConn = conn
	lb.vmNext = rand.Uint32()

	// reset degradation when vmConn changes
	lb.degradeConn = make([]uint32, len(conn))
	lb.numDegradConn = 0
}

// func (lb *ClientLB) updateServerlessConn(conn []*grpc.ClientConn) {
// 	lb.mu.Lock()
// 	defer lb.mu.Unlock()

// 	go closeConn(lb.serverlessConn)
// 	lb.serverlessConn = conn
// 	lb.serverlessNext = rand.Uint32()
// }

func (lb *ClientLB) watchService() {
	if !lb.enable {
		return
	}

	// use knative service as serverless
	serverlessConn, err := utils.NewConn(lb.serverless + ":80")
	if err != nil {
		log.Fatalf("Failed to connect to serverless: %v", err)
	}
	lb.serverlessConn = []*grpc.ClientConn{serverlessConn}

	// watch
	vmCh := utils.WatchEndpointConns(lb.kubeClient, lb.vm, lb.port)
	// serverlessCh := utils.WatchEndpointConns(lb.kubeClient, lb.serverless, lb.port)
	for {
		select {
		case vmConn := <-vmCh:
			log.Printf("%s conn: %v", lb.vm, len(vmConn))
			lb.updateVmConn(vmConn)
			// case serverlessConn := <-serverlessCh:
			// 	log.Printf("%s conn: %v", lb.serverless, len(serverlessConn))
			// 	lb.updateServerlessConn(serverlessConn)
		}
	}
}

func (lb *ClientLB) selectConn() (*grpc.ClientConn, int) {
	lb.mu.RLock()
	defer lb.mu.RUnlock()

	// no vm connection
	if len(lb.vmConn) == 0 {
		if len(lb.serverlessConn) == 0 {
			// no available connection
			return nil, -1
		}

		// select serverless connection
		idx := int(atomic.AddUint32(&lb.serverlessNext, 1) % uint32(len(lb.serverlessConn)))
		return lb.serverlessConn[idx], -1
	}

	// no serverless connection
	if len(lb.serverlessConn) == 0 {
		// select vm connection regradless of degradation
		idx := int(atomic.AddUint32(&lb.vmNext, 1) % uint32(len(lb.vmConn)))
		return lb.vmConn[idx], idx
	}

	// portion of request go to serverless
	r := float64(atomic.LoadInt32(&lb.numDegradConn)) / float64(len(lb.vmConn))
	log.Printf("degrade portion: %v", r)
	if rand.Float64() < r {
		log.Printf("select serverless connection")
		idx := int(atomic.AddUint32(&lb.serverlessNext, 1) % uint32(len(lb.serverlessConn)))
		return lb.serverlessConn[idx], -1
	}

	// select vm connection
	for i := 0; i < len(lb.vmConn); i++ {
		idx := int(atomic.AddUint32(&lb.vmNext, 1) % uint32(len(lb.vmConn)))
		if atomic.LoadUint32(&lb.degradeConn[idx]) == 0 {
			log.Printf("select vm connection")
			return lb.vmConn[idx], idx
		}
	}

	// all vm connections are degraded
	idx := int(atomic.AddUint32(&lb.serverlessNext, 1) % uint32(len(lb.serverlessConn)))
	return lb.serverlessConn[idx], -1
}

func (lb *ClientLB) updateConnDegrade(idx int, overload bool) {
	lb.mu.RLock()
	defer lb.mu.RUnlock()

	// serverless connection or not overloaded, no need to degrade
	if idx < 0 || !overload {
		return
	}

	for {
		old := atomic.LoadUint32(&lb.degradeConn[idx])
		if atomic.CompareAndSwapUint32(&lb.degradeConn[idx], old, uint32(lb.degradeInterval)) {
			if old == 0 {
				atomic.AddInt32(&lb.numDegradConn, 1)
			}
			break
		}
	}
}

func (lb *ClientLB) updateDegrade() {
	if !lb.enable {
		return
	}

	ticker := time.NewTicker(UpdateInterval * time.Millisecond)
	for {
		<-ticker.C
		lb.mu.RLock()
		for i := range lb.degradeConn {
			for {
				oldi := atomic.LoadUint32(&lb.degradeConn[i])
				if oldi == 0 {
					break
				}

				newi := oldi
				if newi < UpdateInterval {
					newi = 0
				} else {
					newi -= UpdateInterval
				}
				if atomic.CompareAndSwapUint32(&lb.degradeConn[i], oldi, newi) {
					if newi == 0 {
						atomic.AddInt32(&lb.numDegradConn, -1)
					}
					break
				}
			}
		}
		lb.mu.RUnlock()
	}
}

func (lb *ClientLB) UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		if !lb.enable {
			return invoker(ctx, method, req, reply, cc, opts...)
		}

		conn, idx := lb.selectConn()
		if conn == nil {
			log.Fatal("no available connection")
		}

		var header metadata.MD
		opts = append(opts, grpc.Header(&header))
		err := invoker(ctx, method, req, reply, conn, opts...)

		overload, ok := header[OverloadHeaderKey]
		go lb.updateConnDegrade(idx, ok && overload[0] == "true")

		return err
	}
}
