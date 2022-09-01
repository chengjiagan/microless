package usermentionserver

import (
	"microless/socialnetwork/proto/user"
	pb "microless/socialnetwork/proto/usermention"
	"microless/socialnetwork/utils"

	"go.uber.org/zap"
)

type UserMentionService struct {
	pb.UnimplementedUserMentionServiceServer
	logger     *zap.SugaredLogger
	userClient user.UserServiceClient
}

func NewServer(logger *zap.SugaredLogger, config *utils.Config) (*UserMentionService, error) {
	conn, err := utils.NewConn(config.Service.User)
	if err != nil {
		return nil, err
	}

	return &UserMentionService{
		logger:     logger,
		userClient: user.NewUserServiceClient(conn),
	}, nil
}
