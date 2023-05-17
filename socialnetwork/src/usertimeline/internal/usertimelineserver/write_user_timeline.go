package usertimelineserver

import (
	"context"

	pb "microless/socialnetwork/proto/usertimeline"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
				"$each":     bson.A{postOid},
				"$position": 0,
			},
		},
	}
	res, err := s.mongodb.UpdateOne(ctx, query, update)
	if err != nil {
		s.logger.Errorw("Failed to update user timeline", "user_id", req.UserId, "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	if res.MatchedCount < 1 {
		s.logger.Errorw("Unknown user", "user_id", req.UserId)
		return nil, status.Errorf(codes.NotFound, "user_id: %v doesn't exist", req.UserId)
	}

	// Update user's timeline in redis
	ts := float64(postOid.Timestamp().Unix())
	err = s.updateRedis(ctx, req.UserId, ts, req.PostId)
	if err != nil {
		s.logger.Errorw("Failed to update user timeline in redis", "err", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *UserTimelineService) updateRedis(ctx context.Context, userId string, ts float64, postId string) error {
	redisResult := s.rdb.Exists(ctx, userId)
	if err := redisResult.Err(); err != nil {
		return err
	}

	if redisResult.Val() != 0 {
		err := s.rdb.ZAdd(ctx, userId, &redis.Z{
			Score:  ts,
			Member: postId,
		}).Err()
		if err != nil {
			return err
		}
	}
	return nil
}
