package pageserver

import (
	"context"
	"microless/media/proto"
	"microless/media/proto/castinfo"
	"microless/media/proto/movieinfo"
	"microless/media/proto/moviereview"
	pb "microless/media/proto/page"
	"microless/media/proto/plot"

	"golang.org/x/sync/errgroup"
)

func (s *PageService) ReadPage(ctx context.Context, req *pb.ReadPageRequest) (*pb.ReadPageRespond, error) {
	g, ctx := errgroup.WithContext(ctx)

	// get movie reviews from movie-review-service
	var reviews []*proto.Review
	g.Go(func() error {
		s.logger.Info("Get movie reviews from movie-review-service")
		reviewReq := &moviereview.ReadMovieReviewsRequest{
			MovieId: req.MovieId,
			Start:   req.ReviewStart,
			Stop:    req.ReviewStop,
		}
		reviewResp, err := s.moviereviewClient.ReadMovieReviews(ctx, reviewReq)
		if err != nil {
			s.logger.Errorw("Failed to get movie reviews from movie-review-service", "err", err)
			return err
		}
		reviews = reviewResp.Reviews
		return nil
	})

	// get movie info from movie-info-service
	s.logger.Info("Get movie info from movie-info-service")
	infoReq := &movieinfo.ReadMovieInfoRequest{MovieId: req.MovieId}
	info, err := s.movieinfoClient.ReadMovieInfo(ctx, infoReq)
	if err != nil {
		s.logger.Errorw("Failed to get movie info from movie-info-service", "err", err)
		return nil, err
	}

	// get cast info from cast-info-service
	var castInfos []*proto.CastInfo
	g.Go(func() error {
		s.logger.Info("Get cast infos from cast-info-service")

		castIds := make([]string, len(info.Casts))
		for i, cast := range info.Casts {
			castIds[i] = cast.CastInfoId
		}
		castReq := &castinfo.ReadCastInfoRequest{CastIds: castIds}

		castResp, err := s.castinfoClient.ReadCastInfo(ctx, castReq)
		if err != nil {
			s.logger.Errorw("Failed to get cast infos from cast-info-service", "err", err)
			return err
		}

		castInfos = castResp.CastInfos
		return nil
	})

	// get plot from plot-service
	var moviePlot string
	g.Go(func() error {
		s.logger.Info("Get plot from plot-service")
		plotReq := &plot.ReadPlotRequest{PlotId: info.PlotId}
		plotResp, err := s.plotClient.ReadPlot(ctx, plotReq)
		if err != nil {
			s.logger.Errorw("Failed to get plot from plot-service", "err", err)
			return err
		}
		moviePlot = plotResp.Plot
		return nil
	})

	err = g.Wait()
	if err != nil {
		return nil, err
	}
	return &pb.ReadPageRespond{
		MovieInfo: info,
		CastInfos: castInfos,
		Plot:      moviePlot,
		Reviews:   reviews,
	}, nil
}
