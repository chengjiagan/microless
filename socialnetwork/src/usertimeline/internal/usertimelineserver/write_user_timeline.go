package usertimelineserver

import (
	"context"

	pb "microless/socialnetwork/proto/usertimeline"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserTimelineService) WriteUserTimeline(ctx context.Context, req *pb.WriteUserTimelineRequest) (*emptypb.Empty, error) {
	postOid, _ := primitive.ObjectIDFromHex(req.PostId)
	userOid, _ := primitive.ObjectIDFromHex(req.UserId)

	query := bson.M{"user_id": userOid}
	update := bson.M{
		"$push": bson.M{
			"post_ids": bson.M{
				"$each": bson.A{postOid},
				"$sort": -1,
			},
		},
	}
	err := s.mongodb.FindOneAndUpdate(ctx, query, update).Err()
	if err != nil {
		if err != mongo.ErrNoDocuments {
			s.logger.Errorw("Failed to update user timeline", "user_id", req.UserId, "err", err)
			return nil, status.Errorf(codes.Internal, "Mongo Err: %v", err)
		} else {
			s.logger.Errorw("User doesn't exist", "user_id", req.UserId, "err", err)
			return nil, status.Errorf(codes.NotFound, "user_id: %v doesn't exit in MongoDB", req.UserId)
		}
	}

	// Update user's timeline in redis
	err = s.rdb.Del(ctx, req.UserId).Err()
	if err != nil {
		s.logger.Errorw("Failed to delete user timeline in redis", "err", err)
		return nil, status.Errorf(codes.Internal, "Redis Err: %v", err)
	}

	return &emptypb.Empty{}, nil
}
