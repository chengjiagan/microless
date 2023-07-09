package userserver

import (
	pb "microless/media/proto/user"
	"microless/media/proto/userreview"
	"microless/media/utils"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	logger           *zap.SugaredLogger
	mongodb          *mongo.Collection
	rdb              *redis.Client
	userreviewClient userreview.UserReviewServiceClient
	secret           string
}

func NewServer(logger *zap.SugaredLogger, mongodb *mongo.Collection, rdb *redis.Client, config *utils.Config) (*UserService, error) {
	conn, err := utils.NewConn(config.Service.UserReview)
	if err != nil {
		return nil, err
	}
	userreviewClient := userreview.NewUserReviewServiceClient(conn)

	return &UserService{
		logger:           logger,
		mongodb:          mongodb,
		rdb:              rdb,
		userreviewClient: userreviewClient,
		secret:           config.Secret,
	}, nil
}
