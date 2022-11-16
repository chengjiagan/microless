package usertimelineserver

import (
	"context"

	"microless/socialnetwork/proto"
	"microless/socialnetwork/proto/poststorage"

	pb "microless/socialnetwork/proto/usertimeline"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserTimelineService) ReadUserTimeline(ctx context.Context, req *pb.ReadUserTimelineRequest) (*pb.ReadUserTimelineRespond, error) {
	if req.Stop <= req.Start || req.Start < 0 {
		s.logger.Errorw("Invalid arguments", "user_id", req.UserId, "start", req.Start, "stop", req.Stop)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid arguments for ReadUserTimeline")
	}

	// get user timeline from redis
	postIds, err := s.rdb.ZRevRange(ctx, req.UserId, int64(req.Start), int64(req.Stop-1)).Result()
	if err != nil {
		s.logger.Errorw("Failed to get user timeline from Redis", "user_id", req.UserId, "err", err)
		return nil, status.Errorf(codes.Internal, "Redis Err: %v", err)
	}

	// everything in redis
	if len(postIds) == int(req.Stop-req.Start) {
		postReq := &poststorage.ReadPostsRequest{PostIds: postIds}
		postResp, err := s.poststorageClient.ReadPosts(ctx, postReq)
		if err != nil {
			s.logger.Error("Failed to get post from post-storage-service")
			return nil, err
		}
		return &pb.ReadUserTimelineRespond{Posts: postResp.Posts}, nil
	}

	// find in mongodb
	userOid, _ := primitive.ObjectIDFromHex(req.UserId)
	query := bson.M{"user_id": userOid}
	opts := options.FindOne().SetProjection(
		bson.M{
			"post_ids": bson.M{"$slice": bson.A{0, req.Stop}},
		},
	)
	doc := new(UserTimeline)
	err = s.mongodb.FindOne(ctx, query, opts).Decode(doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "user_id: %v doesn't exit in MongoDB", req.UserId)
		} else {
			s.logger.Errorw("Failed to get user timeline from MongoDB", "user_id", req.UserId, "err", err)
			return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
		}
	}
	postOids := doc.PostIds

	g, ctx := errgroup.WithContext(ctx)
	// update redis
	g.Go(func() error {
		redisUpdate := make([]*redis.Z, len(postOids))
		for i, oid := range postOids {
			redisUpdate[i] = &redis.Z{
				Score:  float64(oid.Timestamp().Unix()),
				Member: oid.Hex(),
			}
		}

		_, err := s.rdb.Pipelined(ctx, func(p redis.Pipeliner) error {
			p.Del(ctx, req.UserId)
			p.ZAdd(ctx, req.UserId, redisUpdate...)
			return nil
		})
		if err != nil {
			s.logger.Errorw("Failed to update user timeline in Redis", "err", err)
		}
		return nil
	})

	// get posts from PostStorage
	var posts []*proto.Post
	g.Go(func() error {
		// prevent out of index panic when requesting more posts than we have
		var nPost int
		if len(postOids) < int(req.Stop) {
			nPost = len(postOids) - int(req.Start)
		} else {
			nPost = int(req.Stop) - int(req.Start)
		}

		postIds = make([]string, nPost)
		for i := range postIds {
			postIds[i] = postOids[i+int(req.Start)].Hex()
		}

		postReq := &poststorage.ReadPostsRequest{PostIds: postIds}
		postResp, err := s.poststorageClient.ReadPosts(ctx, postReq)
		if err != nil {
			s.logger.Error("Failed to get post from post-storage-service")
			return err
		}
		posts = postResp.Posts
		return nil
	})

	err = g.Wait()
	if err != nil {
		return nil, err
	}
	return &pb.ReadUserTimelineRespond{Posts: posts}, nil
}
