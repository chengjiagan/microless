package poststorageserver

import (
	"context"
	"encoding/json"

	"microless/socialnetwork/proto"
	pb "microless/socialnetwork/proto/poststorage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *PostStorageService) ReadPosts(ctx context.Context, req *pb.ReadPostsRequest) (*pb.ReadPostsRespond, error) {
	if len(req.PostIds) == 0 {
		return &pb.ReadPostsRespond{}, nil
	}

	// get posts from redis
	posts := make(map[string]*Post, len(req.PostIds))
	postsCache, err := s.rdb.MGet(ctx, req.PostIds...).Result()
	if err != nil {
		s.logger.Warnw("Cannot get post_ids from Redis", "post_ids", req.PostIds, "err", err)
	}
	for i, v := range postsCache {
		if v != nil {
			post := new(Post)
			json.Unmarshal([]byte(v.(string)), post)
			posts[req.PostIds[i]] = post
		}
	}

	// got all posts from redis
	if len(posts) == len(req.PostIds) {
		pbPosts := make([]*proto.Post, len(posts))
		for i, id := range req.PostIds {
			pbPosts[i] = posts[id].toProto()
		}
		return &pb.ReadPostsRespond{Posts: pbPosts}, nil
	}

	// get posts from mongodb
	oids := make([]primitive.ObjectID, 0, len(req.PostIds)-len(posts))
	for _, id := range req.PostIds {
		if _, ok := posts[id]; !ok {
			oid, _ := primitive.ObjectIDFromHex(id)
			oids = append(oids, oid)
		}
	}
	query := bson.M{"_id": bson.M{"$in": oids}}
	cursor, err := s.mongodb.Find(ctx, query)
	if err != nil {
		s.logger.Warnw("Failed to find posts from MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	// decode
	var mongoPosts []*Post
	err = cursor.All(ctx, &mongoPosts)
	if err != nil {
		s.logger.Warnw("Failed to find posts from MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	// update redis
	postsMiss := make([]interface{}, 0, len(mongoPosts)*2)
	for _, post := range mongoPosts {
		id := post.PostOid.Hex()
		posts[id] = post
		postJson, _ := json.Marshal(post)
		postsMiss = append(postsMiss, id, postJson)
	}
	_, err = s.rdb.MSet(ctx, postsMiss...).Result()
	if err != nil {
		s.logger.Warnw("Failed to set post to Redis", "err", err)
	}

	// still unknown post_id exists
	if len(posts) != len(req.PostIds) {
		s.logger.Warnw("Unknown post_id")
		return nil, status.Error(codes.NotFound, "Unknown post_id")
	}

	pbPosts := make([]*proto.Post, len(posts))
	for i, id := range req.PostIds {
		pbPosts[i] = posts[id].toProto()
	}
	return &pb.ReadPostsRespond{Posts: pbPosts}, nil
}
