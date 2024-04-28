package redis_try

import (
	redis "github.com/redis/go-redis/v9"
)

func New() *redis.Client {
	// 初始化一个新的redis client
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "123456",         // redis没密码，没有设置，则留空
		DB:       0,                // 使用默认数据库
	})
}
