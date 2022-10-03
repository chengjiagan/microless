package userserver

import (
	pb "microless/media/proto/user"
	"microless/media/proto/userreview"
	"microless/media/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache"
	"go.uber.org/zap"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	logger           *zap.SugaredLogger
	mongodb          *mongo.Collection
	memcached        *otelmemcache.Client
	userreviewClient userreview.UserReviewServiceClient
	secret           string
}

func NewServer(logger *zap.SugaredLogger, mongodb *mongo.Collection, memcached *otelmemcache.Client, config *utils.Config) (*UserService, error) {
	conn, err := utils.NewConn(config.Service.UserReview)
	if err != nil {
		return nil, err
	}
	userreviewClient := userreview.NewUserReviewServiceClient(conn)

	return &UserService{
		logger:           logger,
		mongodb:          mongodb,
		memcached:        memcached,
		userreviewClient: userreviewClient,
		secret:           config.Secret,
	}, nil
}
