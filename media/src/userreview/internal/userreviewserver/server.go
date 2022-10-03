package userreviewserver

import (
	"microless/media/proto/reviewstorage"
	pb "microless/media/proto/userreview"
	"microless/media/utils"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type UserReviewService struct {
	pb.UnimplementedUserReviewServiceServer
	logger              *zap.SugaredLogger
	mongodb             *mongo.Collection
	rdb                 *redis.Client
	reviewstorageClient reviewstorage.ReviewStorageServiceClient
}

func NewServer(logger *zap.SugaredLogger, mongodb *mongo.Collection, rdb *redis.Client, config *utils.Config) (*UserReviewService, error) {
	conn, err := utils.NewConn(config.Service.ReviewStorage)
	if err != nil {
		return nil, err
	}

	return &UserReviewService{
		reviewstorageClient: reviewstorage.NewReviewStorageServiceClient(conn),
		logger:              logger,
		mongodb:             mongodb,
		rdb:                 rdb,
	}, nil
}
