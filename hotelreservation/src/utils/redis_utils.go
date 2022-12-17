package utils

import (
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func NewRedisClient(addr string) (*redis.Client, error) {
	opt, err := redis.ParseURL(addr)
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(opt)
	rdb.AddHook(redisotel.TracingHook{})
	return rdb, nil
}

func ShutdownRedis(rdb *redis.Client, logger *zap.Logger) {
	if err := rdb.Close(); err != nil {
		logger.Fatal(err.Error())
	}
}
