package urlshortenserver

import (
	"context"
	"microless/socialnetwork/proto"
	pb "microless/socialnetwork/proto/urlshorten"
	"microless/socialnetwork/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const HOSTNAME = "http://short-url/"

func (s *UrlShortenService) ComposeUrls(ctx context.Context, req *pb.ComposeUrlsRequest) (*pb.ComposeUrlsRespond, error) {
	// do nothing if input is empty
	if len(req.Urls) == 0 {
		return &pb.ComposeUrlsRespond{}, nil
	}

	// generate shortened urls
	targetUrls := make([]*proto.Url, len(req.Urls))
	mongoUrls := make([]interface{}, len(req.Urls))
	for i, url := range req.Urls {
		u := &Url{
			ExpandedUrl:  url,
			ShortenedUrl: HOSTNAME + utils.GetRandString(10),
		}
		targetUrls[i] = u.toProto()
		mongoUrls[i] = u
	}

	_, err := s.mongodb.InsertMany(ctx, mongoUrls)
	if err != nil {
		s.logger.Errorw("Failed to insert shortened urls to MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}

	return &pb.ComposeUrlsRespond{Urls: targetUrls}, nil
}
