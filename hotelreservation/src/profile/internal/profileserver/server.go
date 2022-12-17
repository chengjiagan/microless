package profileserver

import (
	"microless/hotelreservation/proto/geo"
	pb "microless/hotelreservation/proto/profile"
	"microless/hotelreservation/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache"
	"go.uber.org/zap"
)

type ProfileService struct {
	pb.UnimplementedProfileServiceServer
	logger    *zap.SugaredLogger
	mongodb   *mongo.Collection
	memcached *otelmemcache.Client
	geoClient geo.GeoServiceClient
}

func NewServer(logger *zap.SugaredLogger, mongodb *mongo.Collection, memcached *otelmemcache.Client, config *utils.Config) (*ProfileService, error) {
	conn, err := utils.NewConn(config.Service.Geo)
	if err != nil {
		return nil, err
	}
	geoClient := geo.NewGeoServiceClient(conn)

	return &ProfileService{
		logger:    logger,
		memcached: memcached,
		mongodb:   mongodb,
		geoClient: geoClient,
	}, nil
}
