package flightsserver

import (
	pb "microless/acmeair/proto/flights"

	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache"
	"go.uber.org/zap"
)

type FlightsService struct {
	pb.UnimplementedFlightsServiceServer
	logger    *zap.SugaredLogger
	mongodb   *mongo.Collection
	memcached *otelmemcache.Client
}

func NewServer(logger *zap.SugaredLogger, mongodb *mongo.Collection, memcached *otelmemcache.Client) (*FlightsService, error) {
	return &FlightsService{
		logger:    logger,
		mongodb:   mongodb,
		memcached: memcached,
	}, nil
}
