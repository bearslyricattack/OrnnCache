package lfu

import (
	"OrnnCache/cache"
	"context"
	"sort"
	"time"
)

// CacheLfu 实现lfu缓存淘汰策略的结构体
type CacheLfu struct {
	BaseClient cache.Client
	//map维护key的访问次数
	keymap map[string]int
	//使用的最大个数
	maxNumbers int
	//当前使用的个数
	numbers int
}

// Set 所有方法均使用全部采用拷贝指针的方式传递对象
func (c *CacheLfu) Set(ctx context.Context, key string, value interface{}, d time.Duration) {
	//set时.把key值加入map
	c.keymap[key] = 0
	c.numbers += 1
	c.BaseClient.Set(ctx, key, value, d)
}

// Get bool标志有没有找到对象
func (c *CacheLfu) Get(ctx context.Context, k string) (interface{}, bool) {
	//当前key每次被访问时，都将其值加一
	c.keymap[k] += 1
	return c.BaseClient.Get(ctx, k)
}

// GetWithTTL bool标志有没有找到对象
func (c *CacheLfu) GetWithTTL(ctx context.Context, k string) (interface{}, bool) {
	//当前key每次被访问时，都将其值加一
	c.keymap[k] += 1
	return c.BaseClient.GetWithTTL(ctx, k)
}

func (c *CacheLfu) Replace(ctx context.Context, k string, x interface{}, d time.Duration) error {
	return c.BaseClient.Replace(ctx, k, x, d)
}

// Delete 从缓存中删除项目。如果key不在缓存中，则不执行任何操作。
func (c *CacheLfu) Delete(ctx context.Context, k string) (interface{}, bool) {
	//同时删除map中的元素
	delete(c.keymap, k)
	c.BaseClient.Delete(ctx, k)
	return nil, false
}

// ItemCount 返回map中元素个数
func (c *CacheLfu) ItemCount(ctx context.Context) int {
	return c.BaseClient.ItemCount(ctx)
}

// Flush 刷新缓存
func (c *CacheLfu) Flush(ctx context.Context) {
	c.BaseClient.Flush(ctx)
}

// DeleteExpired 删除过期缓存
func (c *CacheLfu) DeleteExpired(ctx context.Context) {
	//排序
	c.keymap = sortByInt(c.keymap)
	//删除
	for c.numbers >= c.maxNumbers {
		// 遍历map并返回第一个元素
		for key, _ := range c.keymap {
			delete(c.keymap, key)
			break
		}
		c.numbers--
	}
	return
}

type Pair struct {
	key   string
	value int
}

func sortByInt(inputMap map[string]int) map[string]int {
	// 将map中的键值对放入切片中
	pairs := make([]Pair, 0, len(inputMap))
	for key, value := range inputMap {
		pairs = append(pairs, Pair{key, value})
	}

	// 对切片按照值进行排序
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value > pairs[j].value
	})

	// 构建排序后的map
	sortedMap := make(map[string]int)
	for _, pair := range pairs {
		sortedMap[pair.key] = pair.value
	}

	return sortedMap
}
