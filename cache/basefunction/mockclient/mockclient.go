package mockclient

import (
	"OrnnCache/cache"
	"context"
	"sync"
	"time"
)

// MockClient 实现缓存屏蔽的结构体
type MockClient struct {
	BaseClient cache.Client
	mu         sync.RWMutex
}

func (m *MockClient) Set(ctx context.Context, key string, value interface{}, d time.Duration) {
	if ctx.Value("test") == "1" {
		m.BaseClient.Set(ctx, key, value, d)
	} else {
		return
	}
}

func (m *MockClient) Get(ctx context.Context, k string) (interface{}, bool) {
	if ctx.Value("test") == "1" {
		m.BaseClient.Get(ctx, k)
	}
	return nil, false
}

func (m *MockClient) Replace(ctx context.Context, k string, x interface{}, d time.Duration) error {
	if ctx.Value("test") == "1" {
		err := m.BaseClient.Replace(ctx, k, x, d)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MockClient) Delete(ctx context.Context, k string) (interface{}, bool) {
	if ctx.Value("test") == "1" {
		m.BaseClient.Delete(ctx, k)
	}
	return nil, false
}

func (m *MockClient) ItemCount(ctx context.Context) int {
	if ctx.Value("test") == "1" {
		m.BaseClient.ItemCount(ctx)
	}
	return 0
}

func (m *MockClient) Flush(ctx context.Context) {
	if ctx.Value("test") == "1" {
		m.BaseClient.Flush(ctx)
	}
	return
}

func (m *MockClient) DeleteExpired(ctx context.Context) {
	if ctx.Value("test") == "1" {
		m.BaseClient.DeleteExpired(ctx)
	} else {
		return
	}
}

func (m *MockClient) GetWithTTL(ctx context.Context, k string) (interface{}, bool) {
	if ctx.Value("test") == "1" {
		m.BaseClient.GetWithTTL(ctx, k)
	}
	return nil, false
}
