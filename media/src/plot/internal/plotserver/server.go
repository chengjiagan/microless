package plotserver

import (
	pb "microless/media/proto/plot"

	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache"
	"go.uber.org/zap"
)

type PlotService struct {
	pb.UnimplementedPlotServiceServer
	logger    *zap.SugaredLogger
	memcached *otelmemcache.Client
	mongodb   *mongo.Collection
}

func NewServer(logger *zap.SugaredLogger, memcached *otelmemcache.Client, mongodb *mongo.Collection) (*PlotService, error) {
	return &PlotService{
		logger:    logger,
		memcached: memcached,
		mongodb:   mongodb,
	}, nil
}
