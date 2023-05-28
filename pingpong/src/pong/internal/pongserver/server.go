package pongserver

import (
	"context"
	pb "microless/pingpong/proto/pong"
	"microless/pingpong/utils"

	"google.golang.org/protobuf/types/known/emptypb"
)

type PongService struct {
	pb.UnimplementedPongServiceServer
}

func NewServer(config *utils.Config) (*PongService, error) {
	return &PongService{}, nil
}

func (s *PongService) Pong(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
