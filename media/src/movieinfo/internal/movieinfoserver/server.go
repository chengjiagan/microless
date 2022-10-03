package movieinfoserver

import (
	pb "microless/media/proto/movieinfo"
	"microless/media/proto/moviereview"
	"microless/media/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache"
	"go.uber.org/zap"
)

type MovieInfoServer struct {
	pb.UnimplementedMovieInfoServiceServer
	logger            *zap.SugaredLogger
	memcached         *otelmemcache.Client
	mongodb           *mongo.Collection
	moviereviewClient moviereview.MovieReviewServiceClient
}

func NewServer(logger *zap.SugaredLogger, memcached *otelmemcache.Client, mongodb *mongo.Collection, config *utils.Config) (*MovieInfoServer, error) {
	conn, err := utils.NewConn(config.Service.MovieReview)
	if err != nil {
		return nil, err
	}
	moviereviewClient := moviereview.NewMovieReviewServiceClient(conn)

	return &MovieInfoServer{
		logger:            logger,
		memcached:         memcached,
		mongodb:           mongodb,
		moviereviewClient: moviereviewClient,
	}, nil
}
