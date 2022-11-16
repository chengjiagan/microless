package moviereviewserver

import (
	"context"
	pb "microless/media/proto/moviereview"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *MovieReviewService) UploadMovieReview(ctx context.Context, req *pb.UploadMovieReviewRequest) (*emptypb.Empty, error) {
	movieOid, _ := primitive.ObjectIDFromHex(req.MovieId)
	reviewOid, _ := primitive.ObjectIDFromHex(req.ReviewId)

	// update movie reviews in mongodb
	s.logger.Info("Update movie review in MongoDB")
	query := bson.M{"movie_id": movieOid}
	update := bson.M{
		"$push": bson.M{
			"review_ids": bson.M{
				"$each": bson.A{reviewOid},
				"$sort": -1,
			},
		},
	}
	err := s.mongodb.FindOneAndUpdate(ctx, query, update).Err()
	if err != nil {
		s.logger.Errorw("Failed to update movie reviews", "movie_id", req.MovieId, "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}

	// invalidate cache in redis
	s.logger.Info("Delete movie reviews in Redis")
	err = s.rdb.Del(ctx, req.MovieId).Err()
	if err != nil {
		s.logger.Warnw("Failed to delete movie review in Redis", "err", err)
	}

	return &emptypb.Empty{}, nil
}