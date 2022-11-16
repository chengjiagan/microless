package socialgraphserver

import (
	"context"
	"fmt"

	pb "microless/socialnetwork/proto/socialgraph"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *SocialGraphService) GetFollowers(ctx context.Context, req *pb.GetFollowersRequest) (*pb.GetFollowersRespond, error) {
	key := fmt.Sprintf("%v:followers", req.UserId)
	result, err := s.rdb.SMembers(ctx, key).Result()
	if err != nil {
		s.logger.Warnw("Failed to get followers from Redis", "user_id", req.UserId, "err", err)
	}

	if len(result) > 0 {
		// cache hit
		return &pb.GetFollowersRespond{FollowersId: result}, nil
	}

	// get from mongodb
	s.logger.Debugw("Get followers cache miss", "user_id", req.UserId)
	userOid, _ := primitive.ObjectIDFromHex(req.UserId)
	query := bson.M{"user_id": userOid}
	doc := new(UserSocialGraph)
	err = s.mongodb.FindOne(ctx, query).Decode(doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "user_id: %v doesn't exit in MongoDB", req.UserId)
		} else {
			s.logger.Errorw("Failed to get user social graph from MongoDB", "user_id", req.UserId, "err", err)
			return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
		}
	}

	// update redis
	followers := make([]string, len(doc.Followers))
	for i, oid := range doc.Followers {
		followers[i] = oid.Hex()
	}
	err = s.rdb.SAdd(ctx, key, followers).Err()
	if err != nil {
		s.logger.Warnw("Failed to update social graph in Redis", "err", err)
	}

	return &pb.GetFollowersRespond{FollowersId: followers}, nil
}

func (s *SocialGraphService) GetFollowees(ctx context.Context, req *pb.GetFolloweesRequest) (*pb.GetFolloweesRespond, error) {
	key := fmt.Sprintf("%v:followees", req.UserId)
	result, err := s.rdb.SMembers(ctx, key).Result()
	if err != nil {
		s.logger.Warnw("Failed to get followees from Redis", "user_id", req.UserId, "err", err)
	}

	if len(result) > 0 {
		// cache hit
		return &pb.GetFolloweesRespond{FolloweesId: result}, nil
	}

	// get from mongodb
	s.logger.Debugw("Get followees cache miss", "user_id", req.UserId)
	userOid, _ := primitive.ObjectIDFromHex(req.UserId)
	query := bson.M{"user_id": userOid}
	doc := new(UserSocialGraph)
	err = s.mongodb.FindOne(ctx, query).Decode(doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "user_id: %v doesn't exit in MongoDB", req.UserId)
		} else {
			s.logger.Errorw("Failed to get user social graph from MongoDB", "user_id", req.UserId, "err", err)
			return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
		}
	}

	// update redis
	followees := make([]string, len(doc.Followees))
	for i, oid := range doc.Followees {
		followees[i] = oid.Hex()
	}
	err = s.rdb.SAdd(ctx, key, followees).Err()
	if err != nil {
		s.logger.Warnw("Failed to update social graph in Redis", "err", err)
	}

	return &pb.GetFolloweesRespond{FolloweesId: followees}, nil
}
