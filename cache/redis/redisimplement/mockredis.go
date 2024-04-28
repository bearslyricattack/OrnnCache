package redisimplement

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

// RedisMockClientImpl 这个结构体是redisMockClient 的实现，用来返回空的数据,实现跳过缓存的目的。
type RedisMockClientImpl struct {
}

func (r RedisMockClientImpl) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return &redis.StatusCmd{}
}

func (r RedisMockClientImpl) Get(ctx context.Context, key string) *redis.StringCmd {
	return &redis.StringCmd{}
}

func (r RedisMockClientImpl) TTL(ctx context.Context, key string) *redis.DurationCmd {
	return &redis.DurationCmd{}
}

func (r RedisMockClientImpl) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	return &redis.IntCmd{}
}

func (r RedisMockClientImpl) Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	return &redis.BoolCmd{}
}

func (r RedisMockClientImpl) FlushAll(ctx context.Context) *redis.StatusCmd {
	return &redis.StatusCmd{}
}
