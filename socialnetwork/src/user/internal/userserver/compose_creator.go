package userserver

import (
	"context"

	"microless/socialnetwork/proto"
	pb "microless/socialnetwork/proto/user"
)

func (s *UserService) ComposeCreatorWithUserId(ctx context.Context, req *pb.ComposeCreatorWithUserIdRequest) (*pb.ComposeCreatorWithUserIdRespond, error) {
	creator := &proto.Creator{
		UserId:   req.UserId,
		Username: req.Username,
	}
	return &pb.ComposeCreatorWithUserIdRespond{Creator: creator}, nil
}

func (s *UserService) ComposeCreatorWithUsername(ctx context.Context, req *pb.ComposeCreatorWithUsernameRequest) (*pb.ComposeCreatorWithUsernameRespond, error) {
	userId, err := s.getUserId(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	creator := &proto.Creator{
		UserId:   userId,
		Username: req.Username,
	}
	return &pb.ComposeCreatorWithUsernameRespond{Creator: creator}, nil
}
