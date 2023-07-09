package castinfoserver

import (
	pb "microless/media/proto/castinfo"

	"github.com/go-redis/redis/v8"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type CastInfoService struct {
	pb.UnimplementedCastInfoServiceServer
	logger  *zap.SugaredLogger
	rdb     *redis.Client
	mongodb *mongo.Collection
}

func NewServer(logger *zap.SugaredLogger, rdb *redis.Client, mongodb *mongo.Collection) (*CastInfoService, error) {
	return &CastInfoService{
		logger:  logger,
		rdb:     rdb,
		mongodb: mongodb,
	}, nil
}
