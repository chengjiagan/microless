package reservationserver

import (
	"microless/hotelreservation/proto/profile"
	pb "microless/hotelreservation/proto/reservation"
	"microless/hotelreservation/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache"
	"go.uber.org/zap"
)

type ReservationService struct {
	pb.UnimplementedReservationServiceServer
	logger        *zap.SugaredLogger
	mongodb       *mongo.Collection
	memcached     *otelmemcache.Client
	profileClient profile.ProfileServiceClient
}

func NewServer(logger *zap.SugaredLogger, mongodb *mongo.Collection, memcached *otelmemcache.Client, config *utils.Config) (*ReservationService, error) {
	conn, err := utils.NewConn(config.Service.Profile)
	if err != nil {
		return nil, err
	}
	profileClient := profile.NewProfileServiceClient(conn)

	return &ReservationService{
		logger:        logger,
		mongodb:       mongodb,
		memcached:     memcached,
		profileClient: profileClient,
	}, nil
}
