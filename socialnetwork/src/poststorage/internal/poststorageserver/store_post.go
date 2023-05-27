package poststorageserver

import (
	"context"

	pb "microless/socialnetwork/proto/poststorage"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *PostStorageService) StorePost(ctx context.Context, req *pb.StorePostRequest) (*pb.StorePostRespond, error) {
	if req.Post.Creator == nil {
		s.logger.Error("Received post with empty creator")
		return nil, status.Error(codes.InvalidArgument, "Received post with empty creator")
	}

	result, err := s.mongodb.InsertOne(ctx, postFromProto(req.Post))
	if err != nil {
		s.logger.Warnw("Failed to insert post to MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}

	s.logger.Infow("Insert new post")
	oid := result.InsertedID.(primitive.ObjectID)
	id := oid.Hex()
	ts := timestamppb.New(oid.Timestamp())
	return &pb.StorePostRespond{PostId: id, Timestamp: ts}, nil
}
