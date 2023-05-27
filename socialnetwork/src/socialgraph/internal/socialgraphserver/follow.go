package socialgraphserver

import (
	"context"
	pb "microless/socialnetwork/proto/socialgraph"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *SocialGraphService) Follow(ctx context.Context, req *pb.FollowRequest) (*emptypb.Empty, error) {
	err := s.follow(ctx, req.UserId, req.FolloweeId)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *SocialGraphService) FollowWithUsername(ctx context.Context, req *pb.FollowWithUsernameRequest) (*emptypb.Empty, error) {
	var userId, followeeId string
	g, gCtx := errgroup.WithContext(ctx)

	// get user id
	g.Go(func() error { return s.getUserId(gCtx, req.UserUsername, &userId) })
	g.Go(func() error { return s.getUserId(gCtx, req.FolloweeUsername, &followeeId) })
	err := g.Wait()
	if err != nil {
		return nil, err
	}

	err = s.follow(ctx, userId, followeeId)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *SocialGraphService) follow(ctx context.Context, userId, followeeId string) error {
	userOid, _ := primitive.ObjectIDFromHex(userId)
	followeeOid, _ := primitive.ObjectIDFromHex(followeeId)
	g, ctx := errgroup.WithContext(ctx)

	// update follower->followee in mongodb
	g.Go(func() error {
		query := bson.M{"user_id": userOid}
		update := bson.M{
			"$addToSet": bson.M{
				"followees": followeeOid,
			},
		}

		res, err := s.mongodb.UpdateOne(ctx, query, update)
		if err != nil {
			s.logger.Warnw("Failed to update user social graph in MongoDB", "user_id", userId, "err", err)
			return status.Errorf(codes.Internal, "MongoDB Err: %v", err)
		}
		if res.MatchedCount < 1 {
			s.logger.Warnw("Unknown user", "user_id", userId)
			return status.Errorf(codes.NotFound, "user_id: %v doesn't exist", userId)
		}
		return nil
	})

	// update followee->follower in mongodb
	g.Go(func() error {
		query := bson.M{"user_id": followeeOid}
		update := bson.M{
			"$addToSet": bson.M{
				"followers": userOid,
			},
		}

		res, err := s.mongodb.UpdateOne(ctx, query, update)
		if err != nil {
			s.logger.Warnw("Failed to update user social graph in MongoDB", "user_id", followeeId, "err", err)
			return status.Errorf(codes.Internal, "MongoDB Err: %v", err)
		}
		if res.MatchedCount < 1 {
			s.logger.Warnw("Unknown user", "user_id", followeeId)
			return status.Errorf(codes.NotFound, "user_id: %v doesn't exist", followeeId)
		}
		return nil
	})

	// update redis
	g.Go(func() error {
		_, err := s.rdb.Pipelined(ctx, func(p redis.Pipeliner) error {
			p.Del(ctx, userId)
			p.Del(ctx, followeeId)
			return nil
		})
		if err != nil {
			s.logger.Warnw("Failed to update user social graph in Redis", "err", err)
		}
		return nil
	})

	err := g.Wait()
	if err != nil {
		return err
	}
	return nil
}
