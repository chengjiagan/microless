package movieinfoserver

import (
	pb "microless/media/proto/movieinfo"
	"microless/media/proto/moviereview"
	"microless/media/utils"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type MovieInfoServer struct {
	pb.UnimplementedMovieInfoServiceServer
	logger            *zap.SugaredLogger
	rdb               *redis.Client
	mongodb           *mongo.Collection
	moviereviewClient moviereview.MovieReviewServiceClient
}

func NewServer(logger *zap.SugaredLogger, rdb *redis.Client, mongodb *mongo.Collection, config *utils.Config) (*MovieInfoServer, error) {
	conn, err := utils.NewConn(config.Service.MovieReview)
	if err != nil {
		return nil, err
	}
	moviereviewClient := moviereview.NewMovieReviewServiceClient(conn)

	return &MovieInfoServer{
		logger:            logger,
		rdb:               rdb,
		mongodb:           mongodb,
		moviereviewClient: moviereviewClient,
	}, nil
}
