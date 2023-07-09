package reviewstorageserver

import (
	pb "microless/media/proto/reviewstorage"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type ReviewStorageService struct {
	pb.UnimplementedReviewStorageServiceServer
	logger  *zap.SugaredLogger
	rdb     *redis.Client
	mongodb *mongo.Collection
}

func NewServer(logger *zap.SugaredLogger, rdb *redis.Client, mongodb *mongo.Collection) (*ReviewStorageService, error) {
	return &ReviewStorageService{
		logger:  logger,
		rdb:     rdb,
		mongodb: mongodb,
	}, nil
}
