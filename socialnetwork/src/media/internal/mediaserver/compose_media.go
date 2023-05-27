package mediaserver

import (
	"context"
	"microless/socialnetwork/proto"
	pb "microless/socialnetwork/proto/media"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *MediaService) ComposeMedia(ctx context.Context, req *pb.ComposeMediaRequest) (*pb.ComposeMediaRespond, error) {
	// check arguments
	if len(req.MediaTypes) != len(req.MediaIds) {
		s.logger.Warnw("Invalid arguments: media_types and media_ids don't have equal length")
		return nil, status.Errorf(codes.InvalidArgument, "input media_types and media_ids don't have equal length")
	}

	media := make([]*proto.Media, len(req.MediaTypes))
	for i := range req.MediaTypes {
		media[i] = &proto.Media{
			MediaId:   req.MediaIds[i],
			MediaType: req.MediaTypes[i],
		}
	}
	return &pb.ComposeMediaRespond{Media: media}, nil
}
