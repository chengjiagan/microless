package castinfoserver

import (
	pb "microless/media/proto/castinfo"

	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache"
	"go.uber.org/zap"
)

type CastInfoService struct {
	pb.UnimplementedCastInfoServiceServer
	logger    *zap.SugaredLogger
	memcached *otelmemcache.Client
	mongodb   *mongo.Collection
}

func NewServer(logger *zap.SugaredLogger, memcached *otelmemcache.Client, mongodb *mongo.Collection) (*CastInfoService, error) {
	return &CastInfoService{
		logger:    logger,
		memcached: memcached,
		mongodb:   mongodb,
	}, nil
}
