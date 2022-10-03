package movieinfoserver

import (
	"context"

	pb "microless/media/proto/movieinfo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *MovieInfoServer) UpdateRating(ctx context.Context, req *pb.UpdateRatingRequest) (*emptypb.Empty, error) {
	// find the movie info
	movie, err := s.getMovieInfo(ctx, req.MovieId)
	if err != nil {
		return nil, err
	}

	// update rating in mongodb
	s.logger.Info("Update rating in MongoDB")
	new_avg := (movie.AvgRating*float64(movie.NumRating) + float64(req.SumUncommittedRating)) / (float64(movie.NumRating) + float64(req.NumUncommittedRating))
	new_num := movie.NumRating + req.NumUncommittedRating
	oid, _ := primitive.ObjectIDFromHex(req.MovieId)
	update := bson.M{
		"$set": bson.M{
			"avg_rating": new_avg,
			"num_rating": new_num,
		},
	}
	_, err = s.mongodb.UpdateByID(ctx, oid, update)
	if err != nil {
		s.logger.Errorw("Failed to update rating in MongoDB", "movie_id", req.MovieId, "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}

	// invalidate cache in memcached
	s.logger.Info("Delete cache in Memcached")
	err = s.memcached.WithContext(ctx).Delete(req.MovieId)
	if err != nil {
		s.logger.Errorw("Failed to delete in Memcached", "movie_id", req.MovieId, "err", err)
	}

	return &emptypb.Empty{}, nil
}
