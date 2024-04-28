package mutexclient

import (
	"OrnnCache/cache"
	"context"
	"sync"
	"time"
)

// MutexClient 实现加锁的结构体,持有原始对象
type MutexClient struct {
	BaseClient cache.Client
	mu         sync.RWMutex
}

func (m *MutexClient) Set(ctx context.Context, key string, value interface{}, d time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.BaseClient.Set(ctx, key, value, d)
}

func (m *MutexClient) Get(ctx context.Context, k string) (interface{}, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.BaseClient.Get(ctx, k)
}

func (m *MutexClient) Replace(ctx context.Context, k string, x interface{}, d time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.BaseClient.Replace(ctx, k, x, d)
}

func (m *MutexClient) Delete(ctx context.Context, k string) (interface{}, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.BaseClient.Delete(ctx, k)
}

func (m *MutexClient) ItemCount(ctx context.Context) int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.BaseClient.ItemCount(ctx)
}

func (m *MutexClient) Flush(ctx context.Context) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Flush(ctx)
}

func (m *MutexClient) DeleteExpired(ctx context.Context) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.DeleteExpired(ctx)
}

func (m *MutexClient) GetWithTTL(ctx context.Context, k string) (interface{}, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.BaseClient.GetWithTTL(ctx, k)
}
