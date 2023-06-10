package loadbalancer

import (
	"context"
	"log"
	"microless/loadbalancer/internal/utils"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"
)

type AmoebaLB struct {
	// params from config
	limit   int
	latency time.Duration

	// connections
	vmConn         *grpc.ClientConn
	serverlessConn *grpc.ClientConn

	// rate
	rate    atomic.Int32
	curRate atomic.Int32

	// vm = 0, use serverless; vm = 1, use vm
	vm atomic.Int32
}

func NewAmoebaLB(addr string) (*AmoebaLB, error) {
	config := utils.GetAmoebaConfig()
	if !config.Enable {
		return nil, nil
	}

	service, _ := splitAddr(addr)
	// create connections to vm and serverless
	vm := "kube://" + service + config.VmPostfix
	serverless := service + config.ServerlessPostfix
	vmConn, err := utils.NewConn(vm)
	if err != nil {
		return nil, err
	}
	serverlessConn, err := utils.NewConn(serverless)
	if err != nil {
		return nil, err
	}

	lb := &AmoebaLB{
		limit:          config.Limit,
		latency:        time.Duration(config.SwitchLatency) * time.Second,
		vmConn:         vmConn,
		serverlessConn: serverlessConn,
	}
	go lb.updateLoop()

	return lb, nil
}

func (lb *AmoebaLB) updateLoop() {
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		lb.updateRate()
	}
}

func (lb *AmoebaLB) updateRate() {
	r := lb.curRate.Load()
	lb.curRate.Store(0)

	oldRate := lb.rate.Load()
	newRate := int32(float64(oldRate)*0.5 + float64(r)*0.5)
	lb.rate.Store(newRate)

	if newRate >= int32(lb.limit) && oldRate < int32(lb.limit) {
		log.Printf("load %d, enable vm after %d seconds", newRate, lb.latency/time.Second)
		go lb.enableVm()
	}
	if newRate < int32(lb.limit) && oldRate >= int32(lb.limit) && lb.vm.Load() == 1 {
		log.Printf("load %d, disable vm", newRate)
		lb.vm.Store(0)
	}
}

func (lb *AmoebaLB) enableVm() {
	time.Sleep(lb.latency)
	if lb.rate.Load() >= int32(lb.limit) {
		log.Printf("load %d, enable vm", lb.rate.Load())
		lb.vm.Store(1)
	} else {
		log.Printf("load %d, not enable vm", lb.rate.Load())
	}
}

func (lb *AmoebaLB) UnaryClientInterceptor() grpc.UnaryClientInterceptor {
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

		lb.curRate.Add(1)
		if lb.vm.Load() == 1 {
			return invoker(ctx, method, req, reply, lb.vmConn, opts...)
		} else {
			return invoker(ctx, method, req, reply, lb.serverlessConn, opts...)
		}
	}
}
