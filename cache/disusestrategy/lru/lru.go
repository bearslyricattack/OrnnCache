package lru

import (
	"OrnnCache/cache"
	"context"
	"time"
)

// CacheLru 实现lru缓存淘汰策略的结构体
type CacheLru struct {
	BaseClient cache.Client
	//用string切片模拟一个简单的FIFO队列
	queue []string
	//使用的最大个数
	maxNumbers int
	//当前使用的个数
	numbers int
}

// Size 暂时不实用大小，而是使用个数
//func Size[K comparable, V any](m map[K]V) int64 {
//	var zeroK K
//	var zeroValue V
//	keySize := unsafe.Sizeof(zeroK)
//	valueSize := unsafe.Sizeof(zeroValue)
//	vo := reflect.ValueOf(m)
//	hm := (*hmap)(unsafe.Pointer(vo.Pointer()))
//	bn := 1<<hm.B + uintptr(hm.noverflow)
//	bz := unsafe.Sizeof(bmap{}) + (keySize+valueSize)*bucketCnt
//	return int64(unsafe.Sizeof(hmap{}) + bz*bn)
//}

// Set 所有方法均使用全部采用拷贝指针的方式传递对象
func (c *CacheLru) Set(ctx context.Context, key string, value interface{}, d time.Duration) {
	//set时.把key值加入队列
	c.queue = append(c.queue, key)
	c.numbers += 1
	c.BaseClient.Set(ctx, key, value, d)
}

// Get bool标志有没有找到对象
func (c *CacheLru) Get(ctx context.Context, k string) (interface{}, bool) {
	//当前key每次被访问时，都将其值移动到队列尾
	for i, v := range c.queue {
		if v == k {
			c.queue = append(c.queue[:i], c.queue[i+1:]...) //删除第i位
			c.queue = append(c.queue, v)                    //添加第i位到末尾
		}
	}
	return c.BaseClient.Get(ctx, k)
}

func (c *CacheLru) Replace(ctx context.Context, k string, x interface{}, d time.Duration) error {
	return c.BaseClient.Replace(ctx, k, x, d)
}

// Delete 从缓存中删除项目。如果key不在缓存中，则不执行任何操作。
func (c *CacheLru) Delete(ctx context.Context, k string) (interface{}, bool) {
	//同时删除队列中的元素
	for i, v := range c.queue {
		if v == k {
			c.queue = append(c.queue[:i], c.queue[i+1:]...) //删除第i位
		}
	}
	c.BaseClient.Delete(ctx, k)
	return nil, false
}

// ItemCount 返回map中元素个数
func (c *CacheLru) ItemCount(ctx context.Context) int {
	return c.BaseClient.ItemCount(ctx)
}

// Flush 刷新缓存
func (c *CacheLru) Flush(ctx context.Context) {
	c.BaseClient.Flush(ctx)
}

// DeleteExpired 删除过期缓存，初始不做任何操作
func (c *CacheLru) DeleteExpired(ctx context.Context) {
	for c.numbers >= c.maxNumbers {
		key := c.queue[0]
		c.BaseClient.Delete(ctx, key)
		c.queue = append(c.queue[1:])
		c.numbers -= 1
	}
	return
}
