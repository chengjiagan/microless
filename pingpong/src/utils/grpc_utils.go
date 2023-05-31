package utils

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"microless/loadbalancer"
)

func NewConn(address string) (*grpc.ClientConn, error) {
	lb, err := loadbalancer.NewClientLB(address)
	if err != nil {
		return nil, fmt.Errorf("failed to create client load balancer: %v", err)
	}

	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			lb.UnaryClientInterceptor(),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %v", err)
	}

	return conn, nil
}

func NewGRPCServer() *grpc.Server {
	stats := loadbalancer.NewStats()
	serverLB := loadbalancer.NewServerLB()
	serverlessLB := loadbalancer.NewServerlessLB(stats)
	go stats.StartMetricServer()

	return grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			serverLB.UnaryServerInterceptor(),
			serverlessLB.UnaryServerInterceptor(),
			stats.UnaryServerInterceptor(),
		),
	)
}
