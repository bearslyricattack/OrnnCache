package puncture

import (
	"OrnnCache/cache/remoteservice/docall"
	"context"
)

// remoteClient 调用远程服务接口
type remoteClient[T, K any] struct {
	remote docall.Remote[T, K]
	group  Group[T]
}

// doCallRemote 发送远程请求
// 如果是分布式场景 多个示例缓存数据不同步的情况 可以直接使用分布式锁即可
func (r *remoteClient[T, K]) update(ctx context.Context, method string, urlPath string, params map[string]interface{}, header map[string]string, body map[string]interface{}) (res K, err error) {
	//r.group.Do(method+urlPath, r.remote.Update(ctx, method, urlPath, params, header, body))
	return
}
