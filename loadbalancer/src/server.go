package loadbalancer

import (
	"context"
	"microless/loadbalancer/internal/utils"
	"sync/atomic"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type ServerLB struct {
	// params from config
	reject   bool
	max      int32
	fill     int32
	interval time.Duration

	tokens int32

	stats *serverStats
}

type serverStats struct {
	reg            *prometheus.Registry
	totalRequests  *prometheus.CounterVec
	requestLatency *prometheus.SummaryVec
}

func NewServerLB() *ServerLB {
	config := utils.GetServerConfig()
	if !config.Enable {
		return nil
	}

	lb := &ServerLB{
		tokens:   int32(config.MaxTokens),
		max:      int32(config.MaxTokens),
		fill:     int32(config.TokensPerFill),
		interval: time.Duration(config.FillInterval) * time.Second,
		stats:    newServerStats(),
	}
	go lb.fillTokens()
	go utils.StartMetricServer(config.MetricAddr, lb.stats.reg)

	return lb
}

// TODO: init stats
func newServerStats() *serverStats {
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

	reg := prometheus.NewRegistry()
	prometheus.WrapRegistererWith(
		prometheus.Labels{"type": "vm"},
		reg,
	).MustRegister(total, latency)

	return &serverStats{
		reg:            reg,
		totalRequests:  total,
		requestLatency: latency,
	}
}

func (lb *ServerLB) fillTokens() {
	ticker := time.NewTicker(lb.interval)
	for {
		<-ticker.C

		for {
			oldt := atomic.LoadInt32(&lb.tokens)
			newt := oldt + lb.fill
			if newt > lb.max {
				newt = lb.max
			}

			if atomic.CompareAndSwapInt32(&lb.tokens, oldt, newt) {
				break
			}
		}
	}
}

func (lb *ServerLB) decreaseTokens() bool {
	for {
		oldt := atomic.LoadInt32(&lb.tokens)
		newt := oldt - 1

		if newt < 0 {
			return false
		}

		if atomic.CompareAndSwapInt32(&lb.tokens, oldt, newt) {
			return true
		}
	}
}

func (lb *ServerLB) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if lb == nil {
			return handler(ctx, req)
		}

		overload := !lb.decreaseTokens()
		if overload && lb.reject {
			header := metadata.Pairs(OverloadHeaderKey, "true")
			grpc.SetHeader(ctx, header)
			return nil, status.Error(codes.ResourceExhausted, "Server is overloaded")
		}

		start := time.Now()
		resp, err := handler(ctx, req)
		elapsed := time.Since(start)

		// update stats
		service, method := utils.GetServiceAndMethod(info)
		code := status.Code(err).String()
		lb.stats.totalRequests.WithLabelValues(service, method, code).Inc()
		lb.stats.requestLatency.WithLabelValues(service, method, code).Observe(elapsed.Seconds())

		if overload {
			header := metadata.Pairs(OverloadHeaderKey, "true")
			grpc.SetHeader(ctx, header)
		}
		return resp, err
	}
}
