package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type UniversalClient interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}
