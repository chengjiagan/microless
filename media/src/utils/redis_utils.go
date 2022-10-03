package utils

import (
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
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
