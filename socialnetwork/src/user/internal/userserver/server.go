package userserver

import (
	"microless/socialnetwork/proto/hometimeline"
	"microless/socialnetwork/proto/socialgraph"
	pb "microless/socialnetwork/proto/user"
	"microless/socialnetwork/proto/usertimeline"
	"microless/socialnetwork/utils"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	socialgraphClient  socialgraph.SocialGraphServiceClient
	usertimelineClient usertimeline.UserTimelineServiceClient
	hometimelineClient hometimeline.HomeTimelineServiceClient
	logger             *zap.SugaredLogger
	mongodb            *mongo.Collection
	rdb                *redis.Client
	secret             string
}

func NewServer(logger *zap.SugaredLogger, mongodb *mongo.Collection, rdb *redis.Client, config *utils.Config) (*UserService, error) {
	conn, err := utils.NewConn(config.Service.SocialGraph)
	if err != nil {
		return nil, err
	}
	socialgraphClient := socialgraph.NewSocialGraphServiceClient(conn)

	conn, err = utils.NewConn(config.Service.UserTimeline)
	if err != nil {
		return nil, err
	}
	usertimelineClient := usertimeline.NewUserTimelineServiceClient(conn)

	conn, err = utils.NewConn(config.Service.HomeTimeline)
	if err != nil {
		return nil, err
	}
	hometimelineClinet := hometimeline.NewHomeTimelineServiceClient(conn)

	return &UserService{
		socialgraphClient:  socialgraphClient,
		usertimelineClient: usertimelineClient,
		hometimelineClient: hometimelineClinet,
		logger:             logger,
		mongodb:            mongodb,
		rdb:                rdb,
		secret:             config.Secret,
	}, nil
}
