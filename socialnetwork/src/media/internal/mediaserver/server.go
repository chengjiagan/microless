package mediaserver

import (
	pb "microless/socialnetwork/proto/media"

	"go.uber.org/zap"
)

type MediaService struct {
	pb.UnimplementedMediaServiceServer
	logger *zap.SugaredLogger
}

func NewServer(logger *zap.SugaredLogger) (*MediaService, error) {
	return &MediaService{
		logger: logger,
	}, nil
}
