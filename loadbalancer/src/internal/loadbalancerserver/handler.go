package loadbalancerserver

import (
	"context"
	"net/http"
	"net/http/httputil"

	"github.com/mwitkow/grpc-proxy/proxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type BackendIndex int

const (
	VM BackendIndex = iota
	SERVERLESS
)

func (s *LoadBalancer) GrpcProxyHandler() proxy.StreamDirector {
	return func(ctx context.Context, _ string) (context.Context, *grpc.ClientConn, error) {
		md, _ := metadata.FromIncomingContext(ctx)
		ctx = metadata.NewOutgoingContext(ctx, md.Copy())
		return ctx, s.grpcConn[s.selectBackend()], nil
	}
}

func (s *LoadBalancer) HttpProxyHandler() *httputil.ReverseProxy {
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		host := s.httpConn[s.selectBackend()]
		req.URL.Host = host
		req.Host = host
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}
	return &httputil.ReverseProxy{Director: director}
}

func (s *LoadBalancer) selectBackend() BackendIndex {
	for {
		cur := s.currentRete.Load()
		if cur < s.rateLimit.Load() {
			if s.currentRete.CompareAndSwap(cur, cur+1) {
				return VM
			}
		} else {
			return SERVERLESS
		}
	}
}
