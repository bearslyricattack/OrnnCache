package penetration

import (
	"OrnnCache/cache"
	"context"
	"hash/fnv"
	"time"
)

// CacheBloom 使用布隆过滤器解决缓存穿透问题
type CacheBloom struct {
	BaseClient cache.Client
	bitArray   []bool
	hashList   []BloomFilterHash
}

// BloomFilterHash 布隆过滤器的hash
type BloomFilterHash func(s string) uint32

// RegisterHash 注册hash函数
func (c *CacheBloom) RegisterHash(h ...BloomFilterHash) {
	c.hashList = append(c.hashList, h...)
}

// 使用New32a实现第一个hash
func hash1(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// 使用New32实现第一个hash
func hash2(s string) uint32 {
	h := fnv.New32()
	h.Write([]byte(s))
	return h.Sum32()
}

// Set 所有方法均使用全部采用拷贝指针的方式传递对象
func (c *CacheBloom) Set(ctx context.Context, key string, value interface{}, d time.Duration) {
	bitLen := uint32(len(c.bitArray))
	for _, v := range c.hashList {
		index := v(key) % bitLen
		c.bitArray[index] = true
	}
	c.BaseClient.Set(ctx, key, value, d)
}

// Exist 判断一个元素是否存在
func (c *CacheBloom) Exist(s string) bool {
	var (
		bitLen         = uint32(len(c.bitArray))
		hashLen        = uint32(len(c.hashList))
		count   uint32 = 0
	)

	for _, v := range c.hashList {
		index := v(s) % bitLen
		if c.bitArray[index] {
			count++
		}
	}

	if hashLen == count {
		return true
	}
	return false
}

// Get bool标志有没有找到对象
func (c *CacheBloom) Get(ctx context.Context, k string) (interface{}, bool) {
	//当前key每次被访问时，都使用布隆过滤器判断
	if c.Exist(k) {
		return c.BaseClient.Get(ctx, k)
	}
	return nil, false
}

func (c *CacheBloom) Replace(ctx context.Context, k string, x interface{}, d time.Duration) error {
	return c.BaseClient.Replace(ctx, k, x, d)
}

// Delete 从缓存中删除项目。如果key不在缓存中，则不执行任何操作。
func (c *CacheBloom) Delete(ctx context.Context, k string) (interface{}, bool) {
	c.BaseClient.Delete(ctx, k)
	return nil, false
}

// ItemCount 返回map中元素个数
func (c *CacheBloom) ItemCount(ctx context.Context) int {
	return c.BaseClient.ItemCount(ctx)
}

// Flush 刷新缓存
func (c *CacheBloom) Flush(ctx context.Context) {
	c.BaseClient.Flush(ctx)
}

// DeleteExpired 删除过期缓存，初始不做任何操作
func (c *CacheBloom) DeleteExpired(ctx context.Context) {
	return
}
