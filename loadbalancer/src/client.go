package loadbalancer

import (
	"context"
	"log"
	"math/rand"
	"microless/loadbalancer/internal/utils"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

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
	// protected by mu
	mu                  sync.Mutex
	serverlessAvailable bool
	vmRatio             float64
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
	vm := service + config.VmPostfix + ":" + port
	vmConn, err := utils.NewConn(vm)
	if err != nil {
		return nil, err
	}
	serverless := service + config.ServerlessPostfix + ":" + port
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
		vmRatio:        1.0,
		upstreamRate:   make(map[string]int),
	}
	go lb.updateLoop()
	go lb.watchServerless()

	return lb, nil
}

func (lb *ClientLB) watchServerless() {
	ctx := context.Background()
	pubsub := lb.rdb.Subscribe(ctx, lb.service)
	defer pubsub.Close()

	ch := pubsub.Channel()
	for msg := range ch {
		lb.mu.Lock()
		lb.serverlessAvailable, _ = strconv.ParseBool(msg.Payload)
		log.Printf("%s serverless available: %t", lb.service, lb.serverlessAvailable)
		lb.mu.Unlock()
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
	defer lb.mu.Unlock()

	if len(lb.upstreamRate) == 0 {
		lb.vmRatio = 1.0
		return
	}

	avg := 0.0
	for _, rate := range lb.upstreamRate {
		avg += float64(rate)
	}
	avg /= float64(len(lb.upstreamRate))
	lb.upstreamRate = make(map[string]int)

	if avg < 0.1 {
		// prevent divide by zero
		lb.vmRatio = 1.0
	} else {
		if lb.serverlessAvailable {
			lb.vmRatio *= float64(lb.vmRateLimit) / avg
		} else {
			// when serverless is not available, the previous ratio is not representative,
			// therefore, we do not update ratio based on the previous ratio,
			// but use the current average rate instead
			lb.vmRatio = float64(lb.vmRateLimit) / avg
		}
		if lb.vmRatio > 1.0 {
			lb.vmRatio = 1.0
		}
	}

	log.Printf("update ratio for %s: %f", lb.service, lb.vmRatio)
}

func (lb *ClientLB) selectConnection() selectT {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	if !lb.serverlessAvailable {
		return vm
	}

	r := rand.Float64()
	if r < lb.vmRatio {
		return vm
	} else {
		return serverless
	}
}

func (lb *ClientLB) updateUpstream(upstream string, rate int) {
	lb.mu.Lock()
	defer lb.mu.Unlock()
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
		lb.updateUpstream(upstream, rate)
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
