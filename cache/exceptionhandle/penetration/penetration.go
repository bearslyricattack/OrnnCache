package penetration

import (
	"OrnnCache/cache"
	"context"
	"time"
)

// CacheKey 实现以缓存空key解决缓存穿透问题的解决方案
type CacheKey struct {
	BaseClient cache.Client
	//最大访问次数
	maxNumber int
	//当前访问次数
	numbers int
}

// Set 所有方法均使用全部采用拷贝指针的方式传递对象
func (c *CacheKey) Set(ctx context.Context, key string, value interface{}, d time.Duration) {
	c.BaseClient.Set(ctx, key, value, d)
}

// Get bool标志有没有找到对象
func (c *CacheKey) Get(ctx context.Context, k string) (interface{}, bool) {
	//当前key每次被访问时，都将其值移动到队列尾
	if k == "" {
		c.numbers += 1
	}
	if c.numbers > c.maxNumber {
		c.BaseClient.Set(ctx, "", "", 0)
	}
	return c.BaseClient.Get(ctx, k)
}

func (c *CacheKey) Replace(ctx context.Context, k string, x interface{}, d time.Duration) error {
	return c.BaseClient.Replace(ctx, k, x, d)
}

// Delete 从缓存中删除项目。如果key不在缓存中，则不执行任何操作。
func (c *CacheKey) Delete(ctx context.Context, k string) (interface{}, bool) {
	c.BaseClient.Delete(ctx, k)
	return nil, false
}

// ItemCount 返回map中元素个数
func (c *CacheKey) ItemCount(ctx context.Context) int {
	return c.BaseClient.ItemCount(ctx)
}

// Flush 刷新缓存
func (c *CacheKey) Flush(ctx context.Context) {
	c.BaseClient.Flush(ctx)
}

// DeleteExpired 删除过期缓存，初始不做任何操作
func (c *CacheKey) DeleteExpired(ctx context.Context) {
	return
}
