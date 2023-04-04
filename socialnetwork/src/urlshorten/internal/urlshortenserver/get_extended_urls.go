package urlshortenserver

import (
	"context"
	pb "microless/socialnetwork/proto/urlshorten"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UrlShortenService) GetExtendedUrls(ctx context.Context, req *pb.GetExtendedUrlsRequest) (*pb.GetExtendedUrlsRespond, error) {
	urls := make(map[string]string, len(req.ShortenedUrls))

	// get from redis
	urlsCache, err := s.rdb.MGet(ctx, req.ShortenedUrls...).Result()
	if err != nil {
		s.logger.Warnw("Failed to get extened urls from Redis", "err", err)
	}
	for i, item := range urlsCache {
		if item != nil {
			urls[req.ShortenedUrls[i]] = item.(string)
		}
	}

	// everything in redis
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
	// update redis
	urlsMiss := make([]interface{}, len(mongoResult)*2)
	for _, url := range mongoResult {
		urls[url.ShortenedUrl] = url.ExpandedUrl
		urlsMiss = append(urlsMiss, url.ShortenedUrl, url.ExpandedUrl)
	}
	_, err = s.rdb.MSet(ctx, urlsMiss...).Result()
	if err != nil {
		s.logger.Warnw("Failed to set extened urls to Redis", "err", err)
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
