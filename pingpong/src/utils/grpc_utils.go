package utils

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"microless/loadbalancer"
)

func NewConn(address string) (*grpc.ClientConn, error) {
	// lb := loadbalancer.NewClientLB(address)

	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// grpc.WithChainUnaryInterceptor(
		// lb.UnaryClientInterceptor(),
		// ),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewGRPCServer() *grpc.Server {
	stats := loadbalancer.NewStats()
	// serverLB := loadbalancer.NewServerLB()
	// serverlessLB := loadbalancer.NewServerlessLB(stats)
	go stats.StartMetricServer()

	return grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			// serverLB.UnaryServerInterceptor(),
			// serverlessLB.UnaryServerInterceptor(),
			stats.UnaryServerInterceptor(),
		),
	)
}
