package hometimelineserver

import (
	pb "microless/socialnetwork/proto/hometimeline"
	"microless/socialnetwork/proto/poststorage"
	"microless/socialnetwork/proto/socialgraph"
	"microless/socialnetwork/utils"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type HomeTimelineService struct {
	pb.UnimplementedHomeTimelineServiceServer
	logger            *zap.SugaredLogger
	rdb               *redis.Client
	poststorageClient poststorage.PostStorageServiceClient
	socialgraphClient socialgraph.SocialGraphServiceClient
}

func NewServer(logger *zap.SugaredLogger, rdb *redis.Client, config *utils.Config) (*HomeTimelineService, error) {
	conn, err := utils.NewConn(config.Service.PostStorage)
	if err != nil {
		return nil, err
	}
	poststorageClient := poststorage.NewPostStorageServiceClient(conn)

	conn, err = utils.NewConn(config.Service.SocialGraph)
	if err != nil {
		return nil, err
	}
	socialgraphClient := socialgraph.NewSocialGraphServiceClient(conn)

	return &HomeTimelineService{
		logger:            logger,
		rdb:               rdb,
		poststorageClient: poststorageClient,
		socialgraphClient: socialgraphClient,
	}, nil
}
