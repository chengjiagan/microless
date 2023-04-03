package urlshortenserver

import (
	pb "microless/socialnetwork/proto/urlshorten"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type UrlShortenService struct {
	pb.UnimplementedUrlShortenServiceServer
	logger  *zap.SugaredLogger
	rdb     *redis.Client
	mongodb *mongo.Collection
}

func NewServer(logger *zap.SugaredLogger, rdb *redis.Client, mongodb *mongo.Collection) (*UrlShortenService, error) {
	return &UrlShortenService{
		logger:  logger,
		rdb:     rdb,
		mongodb: mongodb,
	}, nil
}
