package docall

import (
	"OrnnCache/cache/remoteservice/datahandle"
	"OrnnCache/cache/remoteservice/servicediscovery"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

type Remote[T, K any] interface {
	Update(ctx context.Context, method string, urlPath string, params map[string]interface{}, header map[string]string, body map[string]interface{}) (res T, err error)
}

// remoteClient 调用远程服务接口
type remoteClient[T, K any] struct {
	//服务相关的参数
	serviceName      string                           //服务名
	discoveryHandler servicediscovery.DiscoveryHandle //服务发现的接口
	dataHandler      datahandle.DataHandle[T, K]      //数据处理接口
}

// NewRemoteClient T 为远端resp结构体，K 为内部使用的结构体 创建调用远程服务的client
func NewRemoteClient[T, K any](serviceName, method, urlPath string, opts ...Option[T, K]) *remoteClient[T, K] {
	client := &remoteClient[T, K]{
		serviceName:      serviceName,
		discoveryHandler: nil,
		dataHandler:      nil,
	}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

// Option 选项 后续支持可动态扩展
type Option[T, K any] func(client *remoteClient[T, K])

// Update  发送远程请求
func (r *remoteClient[T, K]) Update(ctx context.Context, method string, urlPath string, params map[string]string, header map[string]string, body map[string]interface{}) (res K, err error) {
	//访问服务发现，获取服务的动态地址
	url, err := r.discoveryHandler.ServiceDiscovery(r.serviceName, "", nil)
	//拼接完整的url路径
	urlStr := url + urlPath
	//发送请求
	resp, err := sendHTTPRequest(ctx, method, urlStr, params, header, body)
	//解析请求
	if err != nil {
		fmt.Sprintf("GetVersionUpdateConfig send message err. err: %v", err)
		return res, err
	}
	//判断请求返回值是否正确
	respCode := gjson.Get(resp, "code").String()
	if "200" != respCode {
		fmt.Sprintf("GetVersionUpdateConfig request invalid resp err; respStr : %v", resp)
		return res, err
	}
	//获取data
	data := gjson.Get(resp, "data").Raw
	if len(data) == 0 {
		fmt.Sprintf("GetVersionUpdateConfig empty data err")
	}
	err = jsoniter.UnmarshalFromString(data, &res)
	return res, nil
}

// sendHTTPRequest 发送 HTTP 请求
func sendHTTPRequest(ctx context.Context, method string, url string, params map[string]string, header map[string]string, body map[string]interface{}) (string, error) {
	byteBody, err := json.Marshal(body)
	// 将查询参数拼接到 URL 中
	fullURL := url + "?" + buildQueryParams(params)
	// 创建 HTTP 请求客户端
	client := &http.Client{}
	// 创建 HTTP 请求
	req, err := http.NewRequest(method, fullURL, bytes.NewBuffer(byteBody))
	if err != nil {
		return "", err
	}

	// 设置请求头部信息
	for key, value := range header {
		req.Header.Set(key, value)
	}

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应内容
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 将响应内容转换为字符串并返回
	return string(respBody), nil
}

// buildQueryParams 构建查询参数字符串
func buildQueryParams(params map[string]string) string {
	var queryString string
	for key, value := range params {
		queryString += key + "=" + value + "&"
	}
	return queryString[:len(queryString)-1]
}
