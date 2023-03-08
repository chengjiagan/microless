package urlshortenserver

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"microless/socialnetwork/proto"
	pb "microless/socialnetwork/proto/urlshorten"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	mongoUrls := make([]*Url, len(req.Urls))
	for i, url := range req.Urls {
		u := &Url{
			ExpandedUrl:  url,
			ShortenedUrl: genShortenedUrl(url),
		}
		targetUrls[i] = u.toProto()
		mongoUrls[i] = u
	}

	opts := options.Replace().SetUpsert(true)
	for _, url := range mongoUrls {
		query := bson.M{"shortened_url": url.ShortenedUrl}
		_, err := s.mongodb.ReplaceOne(ctx, query, url, opts)
		if err != nil {
			s.logger.Errorw("Failed to insert shortened urls to MongoDB", "err", err)
			return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
		}
	}

	return &pb.ComposeUrlsRespond{Urls: targetUrls}, nil
}

func genShortenedUrl(url string) string {
	hashed := sha256.Sum256([]byte(url))
	encoded := base64.URLEncoding.EncodeToString(hashed[:])
	return HOSTNAME + encoded[:10]
}
