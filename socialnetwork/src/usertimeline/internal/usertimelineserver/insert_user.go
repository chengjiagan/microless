package usertimelineserver

import (
	"context"

	pb "microless/socialnetwork/proto/usertimeline"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserTimelineService) InsertUser(ctx context.Context, req *pb.InsertUserResquest) (*emptypb.Empty, error) {
	oid, _ := primitive.ObjectIDFromHex(req.UserId)
	_, err := s.mongodb.InsertOne(ctx, &UserTimeline{UserId: oid, PostIds: make([]primitive.ObjectID, 0)})
	if err != nil {
		s.logger.Errorw("Failed to insert user to MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}

	s.logger.Info("Insert new user")
	return &emptypb.Empty{}, nil
}
