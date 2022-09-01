package urlshortenserver

import (
	pb "microless/socialnetwork/proto/urlshorten"

	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache"
	"go.uber.org/zap"
)

type UrlShortenService struct {
	pb.UnimplementedUrlShortenServiceServer
	logger    *zap.SugaredLogger
	memcached *otelmemcache.Client
	mongodb   *mongo.Collection
}

func NewServer(logger *zap.SugaredLogger, memcached *otelmemcache.Client, mongodb *mongo.Collection) (*UrlShortenService, error) {
	return &UrlShortenService{
		logger:    logger,
		memcached: memcached,
		mongodb:   mongodb,
	}, nil
}
