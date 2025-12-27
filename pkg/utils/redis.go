package utils

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	rdb       *redis.Client
	ctx       = context.Background()
	onceRedis sync.Once
)

func InitRedis() *redis.Client {
	opt, err := redis.ParseURL(GetEnv("redis.url"))
	if err != nil {
		panic(err)
	}
	return redis.NewClient(opt)
}

func GetRedisClient() (*redis.Client, context.Context) {
	onceRedis.Do(func() {
		if rdb == nil {
			rdb = InitRedis()
		}
	})
	return rdb, ctx
}
