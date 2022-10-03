package moviereviewserver

import (
	"context"
	pb "microless/media/proto/moviereview"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *MovieReviewService) CreateMovie(ctx context.Context, req *pb.CreateMovieRequest) (*emptypb.Empty, error) {
	s.logger.Info("Create movie")
	movieOid, _ := primitive.ObjectIDFromHex(req.MovieId)
	doc := &MovieReview{
		MovieOid:   movieOid,
		ReviewOids: make([]primitive.ObjectID, 0),
	}
	_, err := s.mongodb.InsertOne(ctx, doc)
	if err != nil {
		s.logger.Errorw("Failed to insert new movie into MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	return &emptypb.Empty{}, nil
}
