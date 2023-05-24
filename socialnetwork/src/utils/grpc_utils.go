package utils

import (
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"microless/loadbalancer"
)

func NewConn(address string) (*grpc.ClientConn, error) {
	lb := loadbalancer.NewClientLB(address)

	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			lb.UnaryClientInterceptor(),
			otelgrpc.UnaryClientInterceptor(),
		),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewGRPCServer() *grpc.Server {
	serverLB := loadbalancer.NewServerLB()
	serverlessLB := loadbalancer.NewServerlessLB()

	return grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			serverLB.UnaryServerInterceptor(),
			serverlessLB.UnaryServerInterceptor(),
			otelgrpc.UnaryServerInterceptor(),
		),
	)
}
