package pageserver

import (
	"microless/media/proto/castinfo"
	"microless/media/proto/movieinfo"
	"microless/media/proto/moviereview"
	pb "microless/media/proto/page"
	"microless/media/proto/plot"
	"microless/media/utils"

	"go.uber.org/zap"
)

type PageService struct {
	pb.UnimplementedPageServiceServer
	logger            *zap.SugaredLogger
	moviereviewClient moviereview.MovieReviewServiceClient
	movieinfoClient   movieinfo.MovieInfoServiceClient
	castinfoClient    castinfo.CastInfoServiceClient
	plotClient        plot.PlotServiceClient
}

func NewServer(logger *zap.SugaredLogger, config *utils.Config) (*PageService, error) {
	conn, err := utils.NewConn(config.Service.MovieReview)
	if err != nil {
		return nil, err
	}
	moviereviewClient := moviereview.NewMovieReviewServiceClient(conn)

	conn, err = utils.NewConn(config.Service.MovieInfo)
	if err != nil {
		return nil, err
	}
	movieinfoClient := movieinfo.NewMovieInfoServiceClient(conn)

	conn, err = utils.NewConn(config.Service.CastInfo)
	if err != nil {
		return nil, err
	}
	castinfoClient := castinfo.NewCastInfoServiceClient(conn)

	conn, err = utils.NewConn(config.Service.Plot)
	if err != nil {
		return nil, err
	}
	plotClient := plot.NewPlotServiceClient(conn)

	return &PageService{
		logger:            logger,
		moviereviewClient: moviereviewClient,
		movieinfoClient:   movieinfoClient,
		castinfoClient:    castinfoClient,
		plotClient:        plotClient,
	}, nil
}
