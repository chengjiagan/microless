package customerserver

import (
	pb "microless/acmeair/proto/customer"

	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache"
	"go.uber.org/zap"
)

type CustomerService struct {
	pb.UnimplementedCustomerServiceServer
	logger    *zap.SugaredLogger
	mongodb   *mongo.Collection
	memcached *otelmemcache.Client
}

func NewServer(logger *zap.SugaredLogger, mongodb *mongo.Collection, memcached *otelmemcache.Client) (*CustomerService, error) {
	return &CustomerService{
		logger:    logger,
		mongodb:   mongodb,
		memcached: memcached,
	}, nil
}
