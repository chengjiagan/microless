package textserver

import (
	pb "microless/socialnetwork/proto/text"
	"microless/socialnetwork/proto/urlshorten"
	"microless/socialnetwork/proto/usermention"
	"microless/socialnetwork/utils"

	"go.uber.org/zap"
)

type TextService struct {
	pb.UnimplementedTextServiceServer
	logger            *zap.SugaredLogger
	urlshortenClient  urlshorten.UrlShortenServiceClient
	usermentionClient usermention.UserMentionServiceClient
}

func NewServer(logger *zap.SugaredLogger, config *utils.Config) (*TextService, error) {
	conn, err := utils.NewConn(config.Service.UrlShorten)
	if err != nil {
		return nil, err
	}
	urlshortenClient := urlshorten.NewUrlShortenServiceClient(conn)

	conn, err = utils.NewConn(config.Service.UserMention)
	if err != nil {
		return nil, err
	}
	usermentionClient := usermention.NewUserMentionServiceClient(conn)

	return &TextService{
		logger:            logger,
		urlshortenClient:  urlshortenClient,
		usermentionClient: usermentionClient,
	}, nil
}
