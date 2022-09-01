package hometimelineserver

import (
	"context"
	pb "microless/socialnetwork/proto/hometimeline"
	"microless/socialnetwork/proto/socialgraph"

	"github.com/go-redis/redis/v8"
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

	// update redis
	member := &redis.Z{Score: float64(req.Timestamp.Seconds), Member: req.PostId}
	_, err = s.rdb.Pipelined(ctx, func(p redis.Pipeliner) error {
		for _, id := range followers {
			p.ZAddNX(ctx, id, member)
		}
		for _, id := range req.UserMentionsId {
			p.ZAddNX(ctx, id, member)
		}
		return nil
	})
	if err != nil {
		s.logger.Errorw("Failed to update home timeline in Redis", "user_id", req.UserId, "err", err)
		return nil, status.Errorf(codes.Internal, "Redis Err: %v", err)
	}

	return &emptypb.Empty{}, nil
}
