package geoserver

import (
	pb "microless/hotelreservation/proto/geo"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

const (
	maxSearchRadius  = 10 // search radius in km
	maxSearchResults = 5
	queryCacheKey    = "query_locations" // query cache key in redis
)

type GeoService struct {
	pb.UnimplementedGeoServiceServer
	logger  *zap.SugaredLogger
	mongodb *mongo.Collection
	rdb     *redis.Client
}

func NewServer(logger *zap.SugaredLogger, mongodb *mongo.Collection, rdb *redis.Client) (*GeoService, error) {
	return &GeoService{
		logger:  logger,
		mongodb: mongodb,
		rdb:     rdb,
	}, nil
}
