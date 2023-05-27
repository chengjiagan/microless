package loadbalancer

import (
	"context"
	"flag"
	"microless/loadbalancer/internal/utils"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var addr = flag.String("stats_addr", os.Getenv("STATS_ADDR"), "The address of stats server")

type Stats struct {
	reg            *prometheus.Registry
	totalRequests  *prometheus.CounterVec
	requestLatency *prometheus.SummaryVec
}

func NewStats() *Stats {
	config := utils.GetConfig()
	var typeLabel string
	if config.Serverless.Enable {
		typeLabel = "serverless"
	} else {
		typeLabel = "server"
	}

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
		prometheus.Labels{"type": typeLabel},
		reg,
	).MustRegister(total, latency)

	return &Stats{
		reg:            reg,
		totalRequests:  total,
		requestLatency: latency,
	}
}

func (s *Stats) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		start := time.Now()
		resp, err = handler(ctx, req)
		elapsed := time.Since(start)

		service, method := utils.GetServiceAndMethod(info)
		code := status.Code(err).String()
		s.totalRequests.WithLabelValues(service, method, code).Inc()
		s.requestLatency.WithLabelValues(service, method, code).Observe(elapsed.Seconds())
		return
	}
}

func (s *Stats) StartMetricServer() {
	go utils.StartMetricServer(*addr, s.reg)
}
