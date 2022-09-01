package usertimelineserver

import (
	"microless/socialnetwork/utils"

	"microless/socialnetwork/proto/poststorage"
	pb "microless/socialnetwork/proto/usertimeline"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type UserTimelineService struct {
	pb.UnimplementedUserTimelineServiceServer
	logger            *zap.SugaredLogger
	mongodb           *mongo.Collection
	rdb               *redis.Client
	poststorageClient poststorage.PostStorageServiceClient
}

func NewServer(logger *zap.SugaredLogger, mongodb *mongo.Collection, rdb *redis.Client, config *utils.Config) (*UserTimelineService, error) {
	conn, err := utils.NewConn(config.Service.PostStorage)
	if err != nil {
		return nil, err
	}

	return &UserTimelineService{
		poststorageClient: poststorage.NewPostStorageServiceClient(conn),
		logger:            logger,
		mongodb:           mongodb,
		rdb:               rdb,
	}, nil
}
