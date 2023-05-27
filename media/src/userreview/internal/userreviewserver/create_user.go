package userreviewserver

import (
	"context"
	pb "microless/media/proto/userreview"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserReviewService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*emptypb.Empty, error) {
	s.logger.Info("Create user")
	userOid, _ := primitive.ObjectIDFromHex(req.UserId)
	doc := &UserReview{
		UserOid:    userOid,
		ReviewOids: make([]primitive.ObjectID, 0),
	}
	_, err := s.mongodb.InsertOne(ctx, doc)
	if err != nil {
		s.logger.Warnw("Failed to insert new user into MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	return &emptypb.Empty{}, nil
}
