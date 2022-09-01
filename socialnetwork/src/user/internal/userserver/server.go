package userserver

import (
	"microless/socialnetwork/proto/socialgraph"
	pb "microless/socialnetwork/proto/user"
	"microless/socialnetwork/proto/usertimeline"
	"microless/socialnetwork/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache"
	"go.uber.org/zap"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	socialgraphClient  socialgraph.SocialGraphServiceClient
	usertimelineClient usertimeline.UserTimelineServiceClient
	logger             *zap.SugaredLogger
	mongodb            *mongo.Collection
	memcached          *otelmemcache.Client
	secret             string
}

func NewServer(logger *zap.SugaredLogger, mongodb *mongo.Collection, memcached *otelmemcache.Client, config *utils.Config) (*UserService, error) {
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

	return &UserService{
		socialgraphClient:  socialgraphClient,
		usertimelineClient: usertimelineClient,
		logger:             logger,
		mongodb:            mongodb,
		memcached:          memcached,
		secret:             config.Secret,
	}, nil
}
