package loadbalancer

import (
	"context"
	"log"
	"math/rand"
	"microless/loadbalancer/internal/utils"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const ratioMax = 1000

type selectT string

const (
	// selectT is used to indicate which connection to select
	vm         selectT = "vm"
	serverless selectT = "serverless"
)

type ClientLB struct {
	// params from config
	updateInterval time.Duration
	retry          int
	vmRateLimit    int

	// params from NewClientLB() args
	service string

	// for local service connection
	local bool
	conn  *grpc.ClientConn

	// used for load balancing
	vmConn         *grpc.ClientConn
	serverlessConn *grpc.ClientConn
	rdb            *redis.Client

	// for rate estimation
	serverlessAvailable atomic.Bool
	vmRatio             atomic.Int32
	mu                  sync.Mutex
	upstreamRate        map[string]int
}

func splitAddr(addr string) (string, string) {
	splits := strings.Split(addr, ":")
	return splits[0], splits[1]
}

func NewClientLB(addr string) (*ClientLB, error) {
	config := utils.GetClientConfig()
	if !config.Enable {
		return nil, nil
	}

	service, port := splitAddr(addr)
	// check if the service is local
	if localPort, ok := config.LocalServices[service]; ok {
		log.Printf("create local connection for %s", service)
		conn, err := utils.NewConn("localhost:" + localPort)
		if err != nil {
			return nil, err
		}

		return &ClientLB{
			local: true,
			conn:  conn,
		}, nil
	}

	// create remote connections
	vm := "kube://" + service + config.VmPostfix + ":" + port
	vmConn, err := utils.NewConn(vm)
	if err != nil {
		return nil, err
	}
	serverless := "kube://" + service + config.ServerlessPostfix + ":" + port
	serverlessConn, err := utils.NewConn(serverless)
	if err != nil {
		return nil, err
	}

	// create redis client
	opt, err := redis.ParseURL(config.RedisAddr)
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(opt)

	lb := &ClientLB{
		updateInterval: time.Duration(config.UpdateInterval) * time.Second,
		retry:          config.Retry,
		vmRateLimit:    config.ServiceRateLimit[service],
		service:        service,
		rdb:            rdb,
		vmConn:         vmConn,
		serverlessConn: serverlessConn,
		upstreamRate:   make(map[string]int),
	}
	lb.vmRatio.Store(ratioMax)

	go lb.updateLoop()
	go lb.watchServerless()

	return lb, nil
}

func (lb *ClientLB) watchServerless() {
	ctx := context.Background()
	available, err := lb.rdb.Get(ctx, lb.service).Bool()
	if err != nil {
		log.Printf("failed to get serverless availability of %s: %v", lb.service, err)
	}
	lb.serverlessAvailable.Store(available)

	pubsub := lb.rdb.Subscribe(ctx, lb.service)
	defer pubsub.Close()

	ch := pubsub.Channel()
	for msg := range ch {
		available, _ := strconv.ParseBool(msg.Payload)
		lb.serverlessAvailable.Store(available)
		log.Printf("%s serverless available: %t", lb.service, available)
	}
}

func (lb *ClientLB) updateLoop() {
	ticker := time.NewTicker(lb.updateInterval)
	for range ticker.C {
		lb.updateRatio()
	}
}

func (lb *ClientLB) updateRatio() {
	lb.mu.Lock()
	m := lb.upstreamRate
	lb.upstreamRate = make(map[string]int)
	lb.mu.Unlock()

	if len(m) == 0 {
		lb.vmRatio.Store(ratioMax)
		return
	}

	avg := 0.0
	for _, rate := range m {
		avg += float64(rate)
	}
	avg /= float64(len(m))

	ratio := float64(lb.vmRatio.Load()) / ratioMax
	if avg < 0.1 {
		// prevent divide by zero
		ratio = 1.0
	} else {
		if lb.serverlessAvailable.Load() {
			ratio *= float64(lb.vmRateLimit) / avg
		} else {
			// when serverless is not available, the previous ratio is not representative,
			// therefore, we do not update ratio based on the previous ratio,
			// but use the current average rate instead
			ratio = float64(lb.vmRateLimit) / avg
		}
		if ratio > 1.0 {
			ratio = 1.0
		}
	}
	lb.vmRatio.Store(int32(ratio * ratioMax))

	log.Printf("update ratio for %s: %f", lb.service, ratio)
}

func (lb *ClientLB) selectConnection() selectT {
	if !lb.serverlessAvailable.Load() {
		return vm
	}

	r := rand.Int31n(ratioMax)
	if r < lb.vmRatio.Load() {
		return vm
	} else {
		return serverless
	}
}

func (lb *ClientLB) updateUpstream(upstream string, rate int, t time.Time) {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	if time.Since(t) > lb.updateInterval {
		// ignore outdated rate
		return
	}

	lb.upstreamRate[upstream] = rate
}

func (lb *ClientLB) UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) (err error) {
		if lb == nil {
			return invoker(ctx, method, req, reply, cc, opts...)
		}

		if lb.local {
			// local mode, no need to load balance
			return invoker(ctx, method, req, reply, lb.conn, opts...)
		}

		sel := lb.selectConnection()
		switch sel {
		case vm:
			err = lb.sendVm(ctx, method, req, reply, invoker, opts)
		case serverless:
			err = lb.sendServerless(ctx, method, req, reply, invoker, opts)
		}

		return
	}
}

func (lb *ClientLB) sendVm(ctx context.Context, method string, req, reply interface{}, invoker grpc.UnaryInvoker, opts []grpc.CallOption) (err error) {
	// send request
	var header metadata.MD
	opts = append(opts, grpc.Header(&header))
	err = invoker(ctx, method, req, reply, lb.vmConn, opts...)

	// update upstream rate
	val, ok := header[ServerHeaderKey]
	if ok && len(val) > 0 {
		upstream := val[0]
		rate, _ := strconv.Atoi(header[RateHeaderKey][0])
		go lb.updateUpstream(upstream, rate, time.Now())
	}
	return
}

func (lb *ClientLB) sendServerless(ctx context.Context, method string, req, reply interface{}, invoker grpc.UnaryInvoker, opts []grpc.CallOption) (err error) {
	for i := 0; i < lb.retry; i++ {
		err = invoker(ctx, method, req, reply, lb.serverlessConn, opts...)
		// serverless connection may return with error ResourceExhausted
		// in this case, we should retry with another connection
		if status.Code(err) != codes.ResourceExhausted {
			return err
		}
	}
	// if all retries failed, we should try to send to vm
	return lb.sendVm(ctx, method, req, reply, invoker, opts)
}
