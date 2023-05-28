package pingserver

import (
	"context"
	"fmt"
	pb "microless/pingpong/proto/ping"
	"microless/pingpong/proto/pong"
	"microless/pingpong/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PingService struct {
	pb.UnimplementedPingServiceServer
	pongClient pong.PongServiceClient
}

func NewServer(config *utils.Config) (*PingService, error) {
	cc, err := utils.NewConn(config.Service.Pong)
	if err != nil {
		return nil, fmt.Errorf("failed to create pong client: %v", err)
	}
	pongClient := pong.NewPongServiceClient(cc)

	return &PingService{pongClient: pongClient}, nil
}

func (s *PingService) Ping(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	_, err := s.pongClient.Pong(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to pong: %v", err)
	}

	return &emptypb.Empty{}, nil
}
