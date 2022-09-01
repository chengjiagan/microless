package composepostserver

import (
	pb "microless/socialnetwork/proto/composepost"
	"microless/socialnetwork/proto/hometimeline"
	"microless/socialnetwork/proto/media"
	"microless/socialnetwork/proto/poststorage"
	"microless/socialnetwork/proto/text"
	"microless/socialnetwork/proto/user"
	"microless/socialnetwork/proto/usertimeline"
	"microless/socialnetwork/utils"

	"go.uber.org/zap"
)

type ComposePostService struct {
	pb.UnimplementedComposePostServiceServer
	logger             *zap.SugaredLogger
	poststorageClient  poststorage.PostStorageServiceClient
	usertimelineClient usertimeline.UserTimelineServiceClient
	userClient         user.UserServiceClient
	mediaClient        media.MediaServiceClient
	textClient         text.TextServiceClient
	hometimelineClient hometimeline.HomeTimelineServiceClient
}

func NewServer(logger *zap.SugaredLogger, config *utils.Config) (*ComposePostService, error) {
	conn, err := utils.NewConn(config.Service.PostStorage)
	if err != nil {
		return nil, err
	}
	poststorageClient := poststorage.NewPostStorageServiceClient(conn)

	conn, err = utils.NewConn(config.Service.UserTimeline)
	if err != nil {
		return nil, err
	}
	usertimelineClient := usertimeline.NewUserTimelineServiceClient(conn)

	conn, err = utils.NewConn(config.Service.User)
	if err != nil {
		return nil, err
	}
	userClient := user.NewUserServiceClient(conn)

	conn, err = utils.NewConn(config.Service.Media)
	if err != nil {
		return nil, err
	}
	mediaClient := media.NewMediaServiceClient(conn)

	conn, err = utils.NewConn(config.Service.Text)
	if err != nil {
		return nil, err
	}
	textClient := text.NewTextServiceClient(conn)

	conn, err = utils.NewConn(config.Service.HomeTimeline)
	if err != nil {
		return nil, err
	}
	hometimelineClient := hometimeline.NewHomeTimelineServiceClient(conn)

	return &ComposePostService{
		logger:             logger,
		poststorageClient:  poststorageClient,
		usertimelineClient: usertimelineClient,
		userClient:         userClient,
		mediaClient:        mediaClient,
		textClient:         textClient,
		hometimelineClient: hometimelineClient,
	}, nil
}
