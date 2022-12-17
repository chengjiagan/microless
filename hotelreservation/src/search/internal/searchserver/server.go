package searchserver

import (
	"microless/hotelreservation/proto/geo"
	"microless/hotelreservation/proto/profile"
	"microless/hotelreservation/proto/rate"
	"microless/hotelreservation/proto/reservation"
	pb "microless/hotelreservation/proto/search"
	"microless/hotelreservation/utils"

	"go.uber.org/zap"
)

type SearchService struct {
	pb.UnimplementedSearchServiceServer
	logger            *zap.SugaredLogger
	geoClient         geo.GeoServiceClient
	rateClient        rate.RateServiceClient
	reservationClient reservation.ReservationServiceClient
	profileClient     profile.ProfileServiceClient
}

func NewServer(logger *zap.SugaredLogger, config *utils.Config) (*SearchService, error) {
	conn, err := utils.NewConn(config.Service.Geo)
	if err != nil {
		return nil, err
	}
	geoClient := geo.NewGeoServiceClient(conn)

	conn, err = utils.NewConn(config.Service.Rate)
	if err != nil {
		return nil, err
	}
	rateClient := rate.NewRateServiceClient(conn)

	conn, err = utils.NewConn(config.Service.Reservation)
	if err != nil {
		return nil, err
	}
	reservationClient := reservation.NewReservationServiceClient(conn)

	conn, err = utils.NewConn(config.Service.Profile)
	if err != nil {
		return nil, err
	}
	profileClient := profile.NewProfileServiceClient(conn)

	return &SearchService{
		logger:            logger,
		geoClient:         geoClient,
		rateClient:        rateClient,
		reservationClient: reservationClient,
		profileClient:     profileClient,
	}, nil
}
