package reviewstorageserver

import (
	pb "microless/media/proto/reviewstorage"

	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache"
	"go.uber.org/zap"
)

type ReviewStorageService struct {
	pb.UnimplementedReviewStorageServiceServer
	logger    *zap.SugaredLogger
	memcached *otelmemcache.Client
	mongodb   *mongo.Collection
}

func NewServer(logger *zap.SugaredLogger, memcached *otelmemcache.Client, mongodb *mongo.Collection) (*ReviewStorageService, error) {
	return &ReviewStorageService{
		logger:    logger,
		memcached: memcached,
		mongodb:   mongodb,
	}, nil
}
