package redis

import "github.com/go-redis/redis/v8"

func New() *redis.Client {
	// 初始化一个新的redis client
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // redis没密码，没有设置，则留空
		DB:       0,                // 使用默认数据库
	})
}
