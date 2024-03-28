package main

import (
	"OrnnCache/cache/baseclient"
	"OrnnCache/cache/mockclient"
	"OrnnCache/cache/mutexclient"
	"context"
)

// 演示如何使用装饰者模式，根据自己的需求组装对象并使用
func main() {
	ctx := context.Background()

	//只使用基本功能
	var base baseclient.BaseClient
	base.Set(ctx, "test", "test", 0)

	//使用读写锁
	var mutex mutexclient.MutexClient
	mutex.BaseClient = &base
	mutex.Set(ctx, "test", "test", 0)

	//使用缓存屏蔽功能
	var mock mockclient.MockClient
	mock.BaseClient = &base
	mock.Set(ctx, "test", "test", 0)

	//同时使用这两个功能
	mock.BaseClient = &mutex
	mock.Set(ctx, "test", "test", 0)
}
