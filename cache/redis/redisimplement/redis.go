package redisimplement

import (
	"OrnnCache/cache/Error"
	"context"
	"errors"
	redis "github.com/redis/go-redis/v9"
	"time"
)

type RedisClientInterface interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	TTL(ctx context.Context, key string) *redis.DurationCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd
	Set(ctx context.Context, key string, values interface{}, expiration time.Duration) *redis.StatusCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	FlushAll(ctx context.Context) *redis.StatusCmd
}

type RedisStore struct {
	RedisTrueClient RedisClientInterface
	RedisMockClient RedisClientInterface
}

func NewRedis(RedisTrueClient RedisClientInterface, RedisMockClient RedisClientInterface) *RedisStore {
	return &RedisStore{
		RedisMockClient: RedisMockClient,
		RedisTrueClient: RedisTrueClient,
	}
}
func (s *RedisStore) Get(ctx context.Context, key string) (interface{}, error) {
	if ctx.Value("test").(string) == "1" {
		object, err := s.RedisMockClient.Get(ctx, key).Result()
		if errors.Is(err, redis.Nil) {
			return nil, Error.NotFoundWithCause(err)
		}
		return object, err
	} else {
		object, err := s.RedisTrueClient.Get(ctx, key).Result()
		if errors.Is(err, redis.Nil) {
			return nil, Error.NotFoundWithCause(err)
		}
		return object, err
	}
}

func (s *RedisStore) GetWithTTL(ctx context.Context, key string) (interface{}, time.Duration, error) {
	if ctx.Value("test").(string) == "1" {
		object, err := s.RedisMockClient.Get(ctx, key).Result()
		if errors.Is(err, redis.Nil) {
			return nil, 0, Error.NotFoundWithCause(err)
		}

		if err != nil {
			return nil, 0, err
		}

		ttl, err := s.RedisMockClient.TTL(ctx, key).Result()
		if err != nil {
			return nil, 0, err
		}
		return object, ttl, err
	} else {
		object, err := s.RedisTrueClient.Get(ctx, key).Result()
		if errors.Is(err, redis.Nil) {
			return nil, 0, Error.NotFoundWithCause(err)
		}

		if err != nil {
			return nil, 0, err
		}

		ttl, err := s.RedisTrueClient.TTL(ctx, key).Result()
		if err != nil {
			return nil, 0, err
		}
		return object, ttl, err
	}

}

func (s *RedisStore) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {

	if ctx.Value("test").(string) == "1" {
		err := s.RedisMockClient.Set(ctx, key, value, expiration).Err()
		if err != nil {
			return err
		}
		return nil
	} else {
		err := s.RedisTrueClient.Set(ctx, key, value, expiration).Err()
		if err != nil {
			return err
		}
		return nil
	}
}

func (s *RedisStore) Delete(ctx context.Context, key string) error {
	if ctx.Value("test").(string) == "1" {
		_, err := s.RedisMockClient.Del(ctx, key).Result()
		return err
	} else {
		_, err := s.RedisTrueClient.Del(ctx, key).Result()
		return err
	}
}

func (s *RedisStore) Clear(ctx context.Context) error {
	if ctx.Value("test").(string) == "1" {
		if err := s.RedisMockClient.FlushAll(ctx).Err(); err != nil {
			return err
		}
		return nil
	} else {
		if err := s.RedisTrueClient.FlushAll(ctx).Err(); err != nil {
			return err
		}
		return nil
	}
}
