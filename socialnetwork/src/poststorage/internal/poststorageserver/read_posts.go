package poststorageserver

import (
	"context"
	"encoding/json"

	"microless/socialnetwork/proto"
	pb "microless/socialnetwork/proto/poststorage"

	"github.com/bradfitz/gomemcache/memcache"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *PostStorageService) ReadPosts(ctx context.Context, req *pb.ReadPostsRequest) (*pb.ReadPostsRespond, error) {
	if len(req.PostIds) == 0 {
		return &pb.ReadPostsRespond{}, nil
	}

	// get posts from memcached
	posts := make(map[string]*Post, len(req.PostIds))
	postsMc, err := s.memcached.WithContext(ctx).GetMulti(req.PostIds)
	if err != nil {
		s.logger.Warnw("Cannot get post_ids from Memcached", "post_ids", req.PostIds, "err", err)
	}
	for k, v := range postsMc {
		post := new(Post)
		json.Unmarshal(v.Value, post)
		posts[k] = post
	}

	// got all posts from memcached
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
		s.logger.Errorw("Failed to find posts from MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	// decode
	var mongoPosts []*Post
	err = cursor.All(ctx, &mongoPosts)
	if err != nil {
		s.logger.Errorw("Failed to find posts from MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	for _, post := range mongoPosts {
		id := post.PostOid.Hex()
		posts[id] = post

		// upload posts to memcached
		postJson, _ := json.Marshal(post)
		err = s.memcached.WithContext(ctx).Set(&memcache.Item{
			Key:   id,
			Value: postJson,
		})
		if err != nil {
			s.logger.Errorw("Failed to set post to Memcached", "err", err)
		}
	}

	// still unknown post_id exists
	if len(posts) != len(req.PostIds) {
		s.logger.Errorw("Unknown post_id")
		return nil, status.Error(codes.NotFound, "Unknown post_id")
	}

	pbPosts := make([]*proto.Post, len(posts))
	for i, id := range req.PostIds {
		pbPosts[i] = posts[id].toProto()
	}
	return &pb.ReadPostsRespond{Posts: pbPosts}, nil
}
