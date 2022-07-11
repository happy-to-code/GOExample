package main

import (
	"fmt"
	. "github.com/NoBugBoy/httpgo/http"
	"net/http"
)

func main() {
	test()
}

func test() {
	req := &Req{}
	body, err := req.Url("https://www.baidu.com").
		Method(http.MethodGet).                       // 请求方式
		Header("user-agent", "Mozilla/5.0...").       // 请求头
		Header("content-type", "application/1.json"). // 请求头可以设置多个
		Timeout(3).                                   // 请求超时时间
		Retry(3).                                     // 请求错误重试次数
		Chunk().                                      // 开启Chunk不会自动关闭response io,需要自己手动读取response body数据并关闭io 参考Test5分块传输
		Params(Query{                                 // 请求参数,所有请求方式通用，如果get参数携带?id=1则优先使用url参数
			"id": 1,
		}).
		// ProxyUrl("192.168.1.1:8080"). // 配置要使用的代理ip
		// ImportProxy(). // 引入配置文件中的代理ip并随机使用
		// Proxy(). // 启用代理模式
		// Build(). // 创建request,一般不需单独调用，使用方法参考Test1压力测试
		Go().  // 发起请求
		Body() // 获取返回值string
	if err != nil {
		panic(err)
	}
	fmt.Println(body)
	request := req.Request // 保留*http.Request对象以便有需要
	fmt.Println(request)
	response := req.Response // 保留*http.Response对象以便有需要
	fmt.Println(response)
	transport := req.TransportSetting() // 操作Transport进行参数调整
	fmt.Println(transport)
}
