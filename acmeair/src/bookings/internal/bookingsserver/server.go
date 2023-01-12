package bookingsserver

import (
	pb "microless/acmeair/proto/bookings"
	"microless/acmeair/proto/customer"
	"microless/acmeair/proto/flights"
	"microless/acmeair/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache"
	"go.uber.org/zap"
)

type BookingsService struct {
	pb.UnimplementedBookingsServiceServer
	logger         *zap.SugaredLogger
	mongodb        *mongo.Collection
	memcached      *otelmemcache.Client
	customerClient customer.CustomerServiceClient
	flightsClient  flights.FlightsServiceClient
}

func NewServer(logger *zap.SugaredLogger, mongodb *mongo.Collection, memcached *otelmemcache.Client, config *utils.Config) (*BookingsService, error) {
	conn, err := utils.NewConn(config.Service.Customer)
	if err != nil {
		return nil, err
	}
	customerClient := customer.NewCustomerServiceClient(conn)

	conn, err = utils.NewConn(config.Service.Flights)
	if err != nil {
		return nil, err
	}
	flightsClient := flights.NewFlightsServiceClient(conn)

	return &BookingsService{
		logger:         logger,
		mongodb:        mongodb,
		memcached:      memcached,
		customerClient: customerClient,
		flightsClient:  flightsClient,
	}, nil
}
