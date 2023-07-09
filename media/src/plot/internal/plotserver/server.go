package plotserver

import (
	pb "microless/media/proto/plot"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type PlotService struct {
	pb.UnimplementedPlotServiceServer
	logger  *zap.SugaredLogger
	rdb     *redis.Client
	mongodb *mongo.Collection
}

func NewServer(logger *zap.SugaredLogger, rdb *redis.Client, mongodb *mongo.Collection) (*PlotService, error) {
	return &PlotService{
		logger:  logger,
		rdb:     rdb,
		mongodb: mongodb,
	}, nil
}
