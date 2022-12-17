package rateserver

import (
	pb "microless/hotelreservation/proto/rate"

	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache"
	"go.uber.org/zap"
)

type RateService struct {
	pb.UnimplementedRateServiceServer
	logger      *zap.SugaredLogger
	ratePlanDb  *mongo.Collection
	hotelRateDb *mongo.Collection
	memcached   *otelmemcache.Client
}

func NewServer(logger *zap.SugaredLogger, ratePlanDb *mongo.Collection, hotelRateDb *mongo.Collection, memcached *otelmemcache.Client) (*RateService, error) {
	return &RateService{
		logger:      logger,
		memcached:   memcached,
		ratePlanDb:  ratePlanDb,
		hotelRateDb: hotelRateDb,
	}, nil
}
