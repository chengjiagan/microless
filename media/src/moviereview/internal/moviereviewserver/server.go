package moviereviewserver

import (
	pb "microless/media/proto/moviereview"
	"microless/media/proto/reviewstorage"
	"microless/media/utils"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type MovieReviewService struct {
	pb.UnimplementedMovieReviewServiceServer
	logger              *zap.SugaredLogger
	mongodb             *mongo.Collection
	rdb                 *redis.Client
	reviewstorageClient reviewstorage.ReviewStorageServiceClient
}

func NewServer(logger *zap.SugaredLogger, mongodb *mongo.Collection, rdb *redis.Client, config *utils.Config) (*MovieReviewService, error) {
	conn, err := utils.NewConn(config.Service.ReviewStorage)
	if err != nil {
		return nil, err
	}

	return &MovieReviewService{
		reviewstorageClient: reviewstorage.NewReviewStorageServiceClient(conn),
		logger:              logger,
		mongodb:             mongodb,
		rdb:                 rdb,
	}, nil
}
