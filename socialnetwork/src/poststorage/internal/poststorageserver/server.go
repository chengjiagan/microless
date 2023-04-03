package poststorageserver

import (
	pb "microless/socialnetwork/proto/poststorage"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type PostStorageService struct {
	pb.UnimplementedPostStorageServiceServer
	logger  *zap.SugaredLogger
	rdb     *redis.Client
	mongodb *mongo.Collection
}

func NewServer(logger *zap.SugaredLogger, rdb *redis.Client, mongodb *mongo.Collection) (*PostStorageService, error) {
	return &PostStorageService{
		logger:  logger,
		rdb:     rdb,
		mongodb: mongodb,
	}, nil
}
