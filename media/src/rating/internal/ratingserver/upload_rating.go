package ratingserver

import (
	"context"
	"microless/media/proto/movieinfo"
	pb "microless/media/proto/rating"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *RatingService) UploadRating(ctx context.Context, req *pb.UploadRatingRequest) (*emptypb.Empty, error) {
	s.logger.Info("Update rating in movie-info-service")
	infoReq := &movieinfo.UpdateRatingRequest{
		MovieId:              req.MovieId,
		SumUncommittedRating: req.Rating,
		NumUncommittedRating: 1,
	}
	_, err := s.movieinfoClient.UpdateRating(ctx, infoReq)
	if err != nil {
		s.logger.Errorw("Failed to update rating in movie-info-service", "err", err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
