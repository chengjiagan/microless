package loadbalancer

import (
	"context"
	"microless/loadbalancer/internal/utils"
	"sync/atomic"
	"time"

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
	}
	go lb.fillTokens()

	return lb
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
	) (resp interface{}, err error) {
		if lb == nil {
			return handler(ctx, req)
		}

		overload := !lb.decreaseTokens()
		if overload && lb.reject {
			header := metadata.Pairs(OverloadHeaderKey, "true")
			grpc.SetHeader(ctx, header)
			return nil, status.Error(codes.ResourceExhausted, "Server is overloaded")
		}

		resp, err = handler(ctx, req)

		if overload {
			header := metadata.Pairs(OverloadHeaderKey, "true")
			grpc.SetHeader(ctx, header)
		}
		return
	}
}
