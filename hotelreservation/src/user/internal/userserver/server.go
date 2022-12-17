package userserver

import (
	pb "microless/hotelreservation/proto/user"
	"microless/hotelreservation/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache"
	"go.uber.org/zap"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	logger    *zap.SugaredLogger
	mongodb   *mongo.Collection
	memcached *otelmemcache.Client
	secret    string
}

func NewServer(logger *zap.SugaredLogger, mongodb *mongo.Collection, memcached *otelmemcache.Client, config *utils.Config) (*UserService, error) {
	return &UserService{
		logger:    logger,
		mongodb:   mongodb,
		memcached: memcached,
		secret:    config.Secret,
	}, nil
}
