package hometimelineserver

import (
	"context"
	pb "microless/socialnetwork/proto/hometimeline"
	"microless/socialnetwork/proto/socialgraph"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *HomeTimelineService) WriteHomeTimeline(ctx context.Context, req *pb.WriteHomeTimelineRequest) (*emptypb.Empty, error) {
	// get followers from SocialGraph
	userReq := &socialgraph.GetFollowersRequest{UserId: req.UserId}
	userResp, err := s.socialgraphClient.GetFollowers(ctx, userReq)
	if err != nil {
		s.logger.Errorw("Failed to get followers from SocialGraph Service", "err", err)
		return nil, err
	}
	followers := userResp.FollowersId

	// update followers' and mentions' home timeline in mongodb
	postOid, _ := primitive.ObjectIDFromHex(req.PostId)
	userOids := make([]primitive.ObjectID, 0, len(followers)+len(req.UserMentionsId))
	for _, userId := range followers {
		userOid, _ := primitive.ObjectIDFromHex(userId)
		userOids = append(userOids, userOid)
	}
	for _, userId := range req.UserMentionsId {
		userOid, _ := primitive.ObjectIDFromHex(userId)
		userOids = append(userOids, userOid)
	}
	// query and modifier of mongodb
	query := bson.M{
		"user_id": bson.M{
			"$in": userOids,
		},
	}
	update := bson.M{
		"$push": bson.M{
			"post_ids": bson.M{
				"$each": bson.A{postOid},
				"$sort": -1,
			},
		},
	}
	// send request to mongodb
	_, err = s.mongodb.UpdateMany(ctx, query, update)
	if err != nil {
		s.logger.Errorw("Failed to update home timeline", "err", err)
		return nil, status.Errorf(codes.Internal, "Mongo Err: %v", err)
	}

	// delete user's home timeline in redis
	_, err = s.rdb.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Del(ctx, followers...)
		p.Del(ctx, req.UserMentionsId...)
		return nil
	})
	if err != nil {
		s.logger.Warnw("Failed to delete home timeline in redis", "err", err)
	}

	return &emptypb.Empty{}, nil
}
