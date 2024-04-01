package docall

import (
	"OrnnCache/cache/remoteservice/datahandle"
	"OrnnCache/cache/remoteservice/servicediscovery"
	"context"
)

type remoteClient interface {
	update(ctx context.Context, urlPath string, params map[string]interface{}, header map[string]string, body map[string]interface{}) (res string, err error)
}

// CallerClient 调用远程服务接口
type CallerClient[T, K any] struct {
	serviceName      string
	discoveryHandler servicediscovery.DiscoveryHandle
	dataHandler      datahandle.DataHandle[T, K]
}

// doCallRemote 发送远程请求
func (c *CallerClient[T, K]) update(context context.Context, urlPath string, params map[string]interface{}, header map[string]string, body map[string]interface{}) (res K, err error) {
	//访问服务发现，获取服务的动态地址
	//url, err := c.discoveryHandler.ServiceDiscovery(c.serviceName, "", nil)
	//拼接完整的url路径
	//urlStr := url + urlPath
	//发送请求
	//解析请求
	return res, nil
}
