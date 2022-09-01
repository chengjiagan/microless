package userserver

import (
	"context"

	pb "microless/socialnetwork/proto/user"

	"github.com/bradfitz/gomemcache/memcache"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserService) GetUserId(ctx context.Context, req *pb.GetUserIdRequest) (*pb.GetUserIdRespond, error) {
	userId, err := s.getUserId(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserIdRespond{UserId: userId}, nil
}

func (s *UserService) getUserId(ctx context.Context, username string) (string, error) {
	keyMc := username + ":user_id"
	item, err := s.memcached.WithContext(ctx).Get(keyMc)
	if err != nil && err != memcache.ErrCacheMiss {
		s.logger.Errorw("Failed to get from Memcached", "err", err)
	}
	if item != nil {
		s.logger.Debugw("user_id cache hit from memcached", "username", username)
		return string(item.Value), nil
	}

	// cache miss
	result := new(User)
	s.logger.Debugw("user_id cache miss from memcached", "username", username)
	err = s.mongodb.FindOne(ctx, bson.M{"username": username}).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			s.logger.Warnw("User doesn't exist in MongoDB", "username", username)
			return "", status.Errorf(codes.NotFound, "username: %v doesn't exist in MongoDB", username)
		} else {
			s.logger.Errorw("Failed to find user from MongoDB", "err", err)
			return "", status.Errorf(codes.Internal, "MongoDB Err: %v", err)
		}
	}

	s.logger.Debugw("User found in MongoDB", "username", username)
	userId := result.UserOid.Hex()
	err = s.memcached.WithContext(ctx).Set(&memcache.Item{
		Key:   keyMc,
		Value: []byte(userId),
	})
	if err != nil {
		s.logger.Errorw("Failed to set to Memcached", "err", err)
	}

	return userId, nil
}
