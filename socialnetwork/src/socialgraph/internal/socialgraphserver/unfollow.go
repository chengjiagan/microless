package socialgraphserver

import (
	"context"

	pb "microless/socialnetwork/proto/socialgraph"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *SocialGraphService) Unfollow(ctx context.Context, req *pb.UnfollowRequest) (*emptypb.Empty, error) {
	err := s.unfollow(ctx, req.UserId, req.FolloweeId)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *SocialGraphService) UnfollowWithUsername(ctx context.Context, req *pb.UnfollowWithUsernameRequest) (*emptypb.Empty, error) {
	var userId, followeeId string
	g, gCtx := errgroup.WithContext(ctx)

	// get user id
	g.Go(func() error { return s.getUserId(gCtx, req.UserUsername, &userId) })
	g.Go(func() error { return s.getUserId(gCtx, req.FolloweeUsername, &followeeId) })
	err := g.Wait()
	if err != nil {
		return nil, err
	}

	err = s.unfollow(ctx, userId, followeeId)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *SocialGraphService) unfollow(ctx context.Context, userId, followeeId string) error {
	userOid, _ := primitive.ObjectIDFromHex(userId)
	followeeOid, _ := primitive.ObjectIDFromHex(followeeId)
	g, ctx := errgroup.WithContext(ctx)

	// remove follower->followee in mongodb
	g.Go(func() error {
		query := bson.M{"user_id": userOid}
		update := bson.M{
			"$pull": bson.M{
				"followees": followeeOid,
			},
		}
		err := s.mongodb.FindOneAndUpdate(ctx, query, update).Err()
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return status.Errorf(codes.NotFound, "user_id: %v doesn't exist", followeeId)
			} else {
				s.logger.Errorw("Failed to update user social graph in MongoDB", "user_id", userId, "err", err)
				return status.Errorf(codes.Internal, "MongoDB Err: %v", err)
			}
		}
		return nil
	})

	// remove followee->follower in mongodb
	g.Go(func() error {
		query := bson.M{"user_id": followeeOid}
		update := bson.M{
			"$pull": bson.M{
				"followers": userOid,
			},
		}
		err := s.mongodb.FindOneAndUpdate(ctx, query, update).Err()
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return status.Errorf(codes.NotFound, "user_id: %v doesn't exist", followeeId)
			} else {
				s.logger.Errorw("Failed to update user social graph in MongoDB", "user_id", userId, "err", err)
				return status.Errorf(codes.Internal, "MongoDB Err: %v", err)
			}
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
			s.logger.Errorw("Failed to update user social graph in Redis", "err", err)
		}
		return nil
	})

	err := g.Wait()
	if err != nil {
		return err
	}
	return nil
}
