package hometimelineserver

import (
	"context"
	pb "microless/socialnetwork/proto/hometimeline"
	"microless/socialnetwork/proto/poststorage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *HomeTimelineService) ReadHomeTimeline(ctx context.Context, req *pb.ReadHomeTimelineRequest) (*pb.ReadHomeTimelineRespond, error) {
	// check arguments
	if req.Stop <= req.Start || req.Start < 0 {
		s.logger.Errorw("Invalid indices", "user_id", req.UserId, "start", req.Start, "stop", req.Stop)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid arguments for ReadHomeTimeline")
	}

	// get home timeline from redis
	postIds, err := s.rdb.ZRevRange(ctx, req.UserId, int64(req.Start), int64(req.Stop-1)).Result()
	if err != nil {
		s.logger.Errorw("Failed to get home timeline from Redis", "user_id", req.UserId, "err", err)
		return nil, status.Errorf(codes.Internal, "Redis Err: %v", err)
	}

	// get posts from PostStorage
	postReq := &poststorage.ReadPostsRequest{PostIds: postIds}
	postResp, err := s.poststorageClient.ReadPosts(ctx, postReq)
	if err != nil {
		s.logger.Error("Failed to get post from post-storage-service")
		return nil, err
	}

	return &pb.ReadHomeTimelineRespond{Posts: postResp.Posts}, nil
}
