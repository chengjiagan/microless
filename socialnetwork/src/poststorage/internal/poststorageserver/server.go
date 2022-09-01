package poststorageserver

import (
	pb "microless/socialnetwork/proto/poststorage"

	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache"
	"go.uber.org/zap"
)

type PostStorageService struct {
	pb.UnimplementedPostStorageServiceServer
	logger    *zap.SugaredLogger
	memcached *otelmemcache.Client
	mongodb   *mongo.Collection
}

func NewServer(logger *zap.SugaredLogger, memcached *otelmemcache.Client, mongodb *mongo.Collection) (*PostStorageService, error) {
	return &PostStorageService{
		logger:    logger,
		memcached: memcached,
		mongodb:   mongodb,
	}, nil
}
