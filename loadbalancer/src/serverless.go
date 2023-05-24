package loadbalancer

import (
	"context"
	"microless/loadbalancer/internal/queue"
	"microless/loadbalancer/internal/utils"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TODO: implement dynamic concurrency control
type ServerlessLB struct {
	maxConcurrency int
	maxCapacity    int
	stats          *serverlessStats

	mu          sync.Mutex
	concurrency int
	tasks       *queue.TaskQueue
}

type serverlessStats struct {
	reg            *prometheus.Registry
	totalRequests  *prometheus.CounterVec
	requestLatency *prometheus.SummaryVec
}

func NewServerlessLB() *ServerlessLB {
	config := utils.GetServerlessConfig()
	if !config.Enable {
		return nil
	}

	sl := &ServerlessLB{
		maxConcurrency: config.MaxConcurrency,
		maxCapacity:    config.MaxCapacity,
		tasks:          queue.NewTaskQueue(config.MaxCapacity),
	}
	sl.stats = newServerlessStats(sl)
	go utils.StartMetricServer(config.MetricAddr, sl.stats.reg)

	return sl
}

func newServerlessStats(lb *ServerlessLB) *serverlessStats {
	total := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: NameRequestTotal,
			Help: HelpRequestTotal,
		},
		[]string{"grpc_service", "grpc_method", "grpc_code"},
	)
	latency := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       NameRequestLatency,
			Help:       HelpRequestLatency,
			Objectives: map[float64]float64{0.95: 0.005, 0.99: 0.001},
		},
		[]string{"grpc_service", "grpc_method", "grpc_code"},
	)
	tasks := prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: NameServerlessTaskCount,
			Help: HelpServerlessTaskCount,
		},
		func() float64 {
			lb.mu.Lock()
			defer lb.mu.Unlock()
			return float64(lb.tasks.Len() + lb.concurrency)
		},
	)

	reg := prometheus.NewRegistry()
	prometheus.WrapRegistererWith(
		prometheus.Labels{"type": "serverless"},
		reg,
	).MustRegister(total, latency, tasks)

	return &serverlessStats{
		reg:            reg,
		totalRequests:  total,
		requestLatency: latency,
	}
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

		lb.mu.Lock()
		// check if the request can be processed
		if lb.concurrency >= lb.maxConcurrency {
			// check if the queue is full
			if lb.tasks.Len() >= lb.maxCapacity {
				lb.mu.Unlock()
				return nil, status.Error(codes.ResourceExhausted, "Serverless queue is full")
			}

			task := make(queue.Task)
			lb.tasks.Push(task)
			// wait for the previous task to finish
			for lb.concurrency >= lb.maxConcurrency {
				lb.mu.Unlock()
				<-task
				lb.mu.Lock()
			}
			// the previous task has finished
			lb.tasks.Pop()
			close(task)
		}
		lb.concurrency++
		lb.mu.Unlock()

		start := time.Now()
		resp, err = handler(ctx, req)
		elapsed := time.Since(start)

		// inform the next task to start
		lb.mu.Lock()
		lb.concurrency--
		if lb.tasks.Len() > 0 {
			next := lb.tasks.Front()
			next <- struct{}{}
		}
		lb.mu.Unlock()

		// update stats
		service, method := utils.GetServiceAndMethod(info)
		code := status.Code(err).String()
		lb.stats.totalRequests.WithLabelValues(service, method, code).Inc()
		lb.stats.requestLatency.WithLabelValues(service, method, code).Observe(elapsed.Seconds())

		return resp, err
	}
}
