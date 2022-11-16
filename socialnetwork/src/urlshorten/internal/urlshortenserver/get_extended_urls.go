package urlshortenserver

import (
	"context"
	pb "microless/socialnetwork/proto/urlshorten"

	"github.com/bradfitz/gomemcache/memcache"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UrlShortenService) GetExtendedUrls(ctx context.Context, req *pb.GetExtendedUrlsRequest) (*pb.GetExtendedUrlsRespond, error) {
	urls := make(map[string]string, len(req.ShortenedUrls))

	// get from memcached
	mcResp, err := s.memcached.WithContext(ctx).GetMulti(req.ShortenedUrls)
	if err != nil {
		s.logger.Warnw("Failed to get extened urls from Memcached", "err", err)
	}
	for k, item := range mcResp {
		urls[k] = string(item.Value)
	}

	// everything in memcacheds
	if len(req.ShortenedUrls) == len(urls) {
		extUrls := make([]string, len(urls))
		for i, url := range req.ShortenedUrls {
			extUrls[i] = urls[url]
		}
		return &pb.GetExtendedUrlsRespond{Urls: extUrls}, nil
	}

	// get from mongodb
	mongoUrls := make([]string, 0, len(req.ShortenedUrls)-len(urls))
	for _, url := range req.ShortenedUrls {
		if _, ok := urls[url]; !ok {
			mongoUrls = append(mongoUrls, url)
		}
	}
	query := bson.M{"shortened_url": bson.M{"$in": mongoUrls}}
	cursor, err := s.mongodb.Find(ctx, query)
	if err != nil {
		s.logger.Errorw("Failed to get extened urls from MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	var mongoResult []Url
	cursor.All(ctx, &mongoResult)
	for _, url := range mongoResult {
		urls[url.ShortenedUrl] = url.ExpandedUrl

		// update memcached
		item := &memcache.Item{
			Key:   url.ShortenedUrl,
			Value: []byte(url.ExpandedUrl),
		}
		s.memcached.WithContext(ctx).Add(item)
	}

	// still unknown url exists
	if len(urls) != len(req.ShortenedUrls) {
		s.logger.Error("Unknown shortened urls")
		return nil, status.Error(codes.NotFound, "Unknown shortened urls")
	}

	extUrls := make([]string, len(urls))
	for i, url := range req.ShortenedUrls {
		extUrls[i] = urls[url]
	}
	return &pb.GetExtendedUrlsRespond{Urls: extUrls}, nil
}
