package redisimplement

import (
	redistry "OrnnCache/cache/redis"
	"context"
	"fmt"
	"math"
	"testing"
	"time"
)

func BenchmarkRedisSet(b *testing.B) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "test", "0")
	store := NewRedis(redistry.New(), nil)
	for k := 0.; k <= 10; k++ {
		n := int(math.Pow(2, k))
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N*n; i++ {
				key := fmt.Sprintf("test-%d", n)
				value := []byte(fmt.Sprintf("value-%d", n))

				store.Set(ctx, key, value, 1*time.Second)
			}
		})
	}
}

func BenchmarkRedisGet(b *testing.B) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "test", "0")
	store := NewRedis(redistry.New(), nil)
	key := "test"
	value := []byte("value")

	store.Set(ctx, key, value, 1*time.Second)

	for k := 0.; k <= 10; k++ {
		n := int(math.Pow(2, k))
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N*n; i++ {
				_, _ = store.Get(ctx, key)
			}
		})
	}
}
