package redisimplement

import (
	"context"
	"github.com/golang/mock/gomock"
	redis "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRedisGet(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "test", "1")
	client := NewMockRedisClientInterface(ctrl)
	client.EXPECT().Get(ctx, "my-key").Return(&redis.StringCmd{})

	store := NewRedis(nil, client)

	// When
	value, err := store.Get(ctx, "my-key")

	// Then
	assert.Nil(t, err)
	assert.NotNil(t, value)
}

func TestRedisSet(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "test", "1")
	cacheKey := "my-key"
	cacheValue := "my-cache-value"

	client := NewMockRedisClientInterface(ctrl)
	client.EXPECT().Set(ctx, "my-key", cacheValue, 5*time.Second).Return(&redis.StatusCmd{})

	store := NewRedis(nil, client)

	// When
	err := store.Set(ctx, cacheKey, cacheValue, 5*time.Second)

	// Then
	assert.Nil(t, err)
}

func TestRedisDelete(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "test", "1")
	cacheKey := "my-key"

	client := NewMockRedisClientInterface(ctrl)
	client.EXPECT().Del(ctx, "my-key").Return(&redis.IntCmd{})

	store := NewRedis(nil, client)

	// When
	err := store.Delete(ctx, cacheKey)

	// Then
	assert.Nil(t, err)
}

func TestRedisClear(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "test", "1")
	client := NewMockRedisClientInterface(ctrl)
	client.EXPECT().FlushAll(ctx).Return(&redis.StatusCmd{})

	store := NewRedis(nil, client)

	// When
	err := store.Clear(ctx)

	// Then
	assert.Nil(t, err)
}
