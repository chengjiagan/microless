package socialgraphserver

import (
	"context"
	pb "microless/socialnetwork/proto/socialgraph"
	"microless/socialnetwork/proto/user"
	"microless/socialnetwork/utils"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type SocialGraphService struct {
	pb.UnimplementedSocialGraphServiceServer
	logger     *zap.SugaredLogger
	mongodb    *mongo.Collection
	rdb        *redis.Client
	userClient user.UserServiceClient
}

func NewServer(logger *zap.SugaredLogger, mongodb *mongo.Collection, rdb *redis.Client, config *utils.Config) (*SocialGraphService, error) {
	conn, err := utils.NewConn(config.Service.User)
	if err != nil {
		return nil, err
	}

	return &SocialGraphService{
		userClient: user.NewUserServiceClient(conn),
		logger:     logger,
		mongodb:    mongodb,
		rdb:        rdb,
	}, nil
}

func (s *SocialGraphService) getUserId(ctx context.Context, username string, userId *string) error {
	userReq := &user.GetUserIdRequest{Username: username}
	resp, err := s.userClient.GetUserId(ctx, userReq)
	if err != nil {
		s.logger.Warnw("Failed to get user_id from user-service", "username", username, "err", err)
		return err
	}
	*userId = resp.UserId
	return nil
}
