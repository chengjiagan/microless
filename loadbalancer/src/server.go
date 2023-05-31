package loadbalancer

import (
	"context"
	"microless/loadbalancer/internal/utils"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ServerLB struct {
	// params from config
	updateInterval time.Duration
	updateRatio    float64
	// params from env var
	name string

	rate    atomic.Int32
	curRate atomic.Int32
}

func NewServerLB() *ServerLB {
	config := utils.GetServerConfig()
	if !config.Enable {
		return nil
	}

	lb := &ServerLB{
		updateInterval: time.Duration(config.UpdateInterval) * time.Second,
		updateRatio:    config.UpdateRatio,
		name:           os.Getenv("POD_NAME"),
	}
	go lb.update()

	return lb
}

func (lb *ServerLB) update() {
	ticker := time.NewTicker(lb.updateInterval)
	for range ticker.C {
		r := float64(lb.curRate.Load()) / lb.updateInterval.Seconds()
		old := float64(lb.rate.Load())
		new := r*lb.updateRatio + old*(1-lb.updateRatio)
		lb.rate.Store(int32(new))
		lb.curRate.Store(0)
	}
}

func (lb *ServerLB) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		if lb == nil {
			return handler(ctx, req)
		}

		lb.curRate.Add(1)
		resp, err = handler(ctx, req)

		header := metadata.Pairs(
			ServerHeaderKey, lb.name,
			RateHeaderKey, strconv.Itoa(int(lb.rate.Load())),
		)
		grpc.SetHeader(ctx, header)
		return
	}
}
