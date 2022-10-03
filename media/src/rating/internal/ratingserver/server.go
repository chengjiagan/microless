package ratingserver

import (
	"microless/media/proto/movieinfo"
	pb "microless/media/proto/rating"
	"microless/media/utils"

	"go.uber.org/zap"
)

type RatingService struct {
	pb.UnimplementedRatingServiceServer
	logger          *zap.SugaredLogger
	movieinfoClient movieinfo.MovieInfoServiceClient
}

func NewServer(logger *zap.SugaredLogger, config *utils.Config) (*RatingService, error) {
	conn, err := utils.NewConn(config.Service.MovieInfo)
	if err != nil {
		return nil, err
	}
	movieinfoClient := movieinfo.NewMovieInfoServiceClient(conn)

	return &RatingService{
		logger:          logger,
		movieinfoClient: movieinfoClient,
	}, nil
}
