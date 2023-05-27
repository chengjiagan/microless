package loadbalancer

import (
	"context"
	"microless/loadbalancer/internal/queue"
	"microless/loadbalancer/internal/utils"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServerlessLB struct {
	// params from config
	totalResources     int
	maxCapacity        int
	methodRequirements map[string]int

	mu               sync.Mutex
	concurrency      int
	currentResources int
	tasks            *queue.TaskQueue
}

func NewServerlessLB(stats *Stats) *ServerlessLB {
	config := utils.GetServerlessConfig()
	if !config.Enable {
		return nil
	}

	sl := &ServerlessLB{
		totalResources:     config.MaxConcurrency * 100,
		maxCapacity:        config.MaxCapacity,
		methodRequirements: config.MethodReqirements,
		tasks:              queue.NewTaskQueue(config.MaxCapacity),
	}

	stats.reg.MustRegister(
		prometheus.NewGaugeFunc(
			prometheus.GaugeOpts{
				Name: NameServerlessTaskTotal,
				Help: HelpServerlessTaskTotal,
			},
			func() float64 {
				sl.mu.Lock()
				defer sl.mu.Unlock()
				return float64(sl.tasks.Len() + sl.concurrency)
			},
		),
		prometheus.NewGaugeFunc(
			prometheus.GaugeOpts{
				Name: NameServerlessTaskRunning,
				Help: HelpServerlessTaskRunning,
			},
			func() float64 {
				sl.mu.Lock()
				defer sl.mu.Unlock()
				return float64(sl.concurrency)
			},
		),
	)

	return sl
}

func (lb *ServerlessLB) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		if lb == nil {
			return handler(ctx, req)
		}

		_, method := utils.GetServiceAndMethod(info)
		methodRequirement := lb.getMethodRequirement(method)
		err = lb.requestResource(methodRequirement)
		if err != nil {
			return
		}

		resp, err = handler(ctx, req)
		lb.releaseResource(methodRequirement)
		return
	}
}

func (lb *ServerlessLB) getMethodRequirement(method string) int {
	if v, ok := lb.methodRequirements[method]; ok {
		return v
	}
	return 100
}

func (lb *ServerlessLB) requestResource(amount int) error {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	// check if the request can be processed
	if lb.currentResources+amount > lb.totalResources {
		// check if the queue is full
		if lb.tasks.Len() >= lb.maxCapacity {
			return status.Error(codes.ResourceExhausted, "Serverless queue is full")
		}

		task := make(queue.Task)
		lb.tasks.Push(task)
		// wait for the previous task to finish
		for lb.currentResources+amount > lb.totalResources {
			lb.mu.Unlock()
			<-task
			lb.mu.Lock()
		}
		// the previous task has finished
		lb.tasks.Pop()
		close(task)
	}
	lb.currentResources += amount
	lb.concurrency++

	return nil
}

func (lb *ServerlessLB) releaseResource(amount int) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	lb.currentResources -= amount
	lb.concurrency--
	if lb.tasks.Len() > 0 {
		next := lb.tasks.Front()
		next <- struct{}{}
	}
}
