package redisimplement

import (
	"context"

	"github.com/go-redis/redis/v8"
	"time"
)

// RedisClient 这个struct是redis client的实现，所有调用redis的地方，所需要的参数写redis client接口，并在调用的时候传入这个结构体实例
type RedisClient struct {
	RedisTrueClient redis.UniversalClient
	RedisMockClient redis.UniversalClient
}

func (r RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	if ctx.Value("test").(string) == "1" {
		return r.RedisMockClient.Set(ctx, key, value, expiration)
	} else {
		return r.RedisTrueClient.Set(ctx, key, value, expiration)
	}
}

func (r RedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	if ctx.Value("test").(string) == "1" {
		return r.RedisMockClient.Get(ctx, key)
	} else {
		return r.RedisTrueClient.Get(ctx, key)
	}
}
