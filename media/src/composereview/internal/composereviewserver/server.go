package composereviewserver

import (
	pb "microless/media/proto/composereview"
	"microless/media/proto/moviereview"
	"microless/media/proto/rating"
	"microless/media/proto/reviewstorage"
	"microless/media/proto/userreview"
	"microless/media/utils"

	"go.uber.org/zap"
)

type ComposeReviewService struct {
	pb.UnimplementedComposeReviewServer
	logger              *zap.SugaredLogger
	reviewstorageClient reviewstorage.ReviewStorageServiceClient
	userreviewClient    userreview.UserReviewServiceClient
	moviereviewClient   moviereview.MovieReviewServiceClient
	ratingClient        rating.RatingServiceClient
}

func NewServer(logger *zap.SugaredLogger, config *utils.Config) (*ComposeReviewService, error) {
	conn, err := utils.NewConn(config.Service.ReviewStorage)
	if err != nil {
		return nil, err
	}
	reviewstorageClient := reviewstorage.NewReviewStorageServiceClient(conn)

	conn, err = utils.NewConn(config.Service.UserReview)
	if err != nil {
		return nil, err
	}
	userreviewclient := userreview.NewUserReviewServiceClient(conn)

	conn, err = utils.NewConn(config.Service.MovieReview)
	if err != nil {
		return nil, err
	}
	moviereviewClient := moviereview.NewMovieReviewServiceClient(conn)

	conn, err = utils.NewConn(config.Service.Rating)
	if err != nil {
		return nil, err
	}
	ratingClient := rating.NewRatingServiceClient(conn)

	return &ComposeReviewService{
		logger:              logger,
		reviewstorageClient: reviewstorageClient,
		userreviewClient:    userreviewclient,
		moviereviewClient:   moviereviewClient,
		ratingClient:        ratingClient,
	}, nil
}
