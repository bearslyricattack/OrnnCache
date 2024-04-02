package avalanche

import (
	"OrnnCache/cache"
	"context"
	"math/rand"
	"sync"
	"time"
)

// 解决缓存雪崩 给对象设置一个随机，但是比较接近的过期时间。

func randomTime() time.Duration {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())
	// 生成0到300秒（五分钟）之间的随机时间
	randomSeconds := rand.Intn(301)
	// 将秒数转换为Duration类型
	randomDuration := time.Duration(randomSeconds) * time.Second
	return randomDuration
}

// avalancheClient 实现缓存屏蔽的结构体
type avalancheClient struct {
	BaseClient cache.Client
	mu         sync.RWMutex
}

func (m *avalancheClient) Set(ctx context.Context, key string, value interface{}, d time.Duration) {
	//加入随机的过期时间
	if ctx.Value("test") == "1" {
		m.BaseClient.Set(ctx, key, value, d+randomTime())
	} else {
		return
	}
}

func (m *avalancheClient) Get(ctx context.Context, k string) (interface{}, bool) {
	if ctx.Value("test") == "1" {
		m.BaseClient.Get(ctx, k)
	}
	return nil, false
}

func (m *avalancheClient) Replace(ctx context.Context, k string, x interface{}, d time.Duration) error {
	if ctx.Value("test") == "1" {
		err := m.BaseClient.Replace(ctx, k, x, d)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *avalancheClient) Delete(ctx context.Context, k string) (interface{}, bool) {
	if ctx.Value("test") == "1" {
		m.BaseClient.Delete(ctx, k)
	}
	return nil, false
}

func (m *avalancheClient) ItemCount(ctx context.Context) int {
	if ctx.Value("test") == "1" {
		m.BaseClient.ItemCount(ctx)
	}
	return 0
}

func (m *avalancheClient) Flush(ctx context.Context) {
	if ctx.Value("test") == "1" {
		m.BaseClient.Flush(ctx)
	}
	return
}

func (m *avalancheClient) DeleteExpired(ctx context.Context) {
	if ctx.Value("test") == "1" {
		m.BaseClient.DeleteExpired(ctx)
	} else {
		return
	}
}
