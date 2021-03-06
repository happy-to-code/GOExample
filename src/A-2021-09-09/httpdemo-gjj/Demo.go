package main

import (
	"flag"
	"fmt"
	"github.com/siddontang/go/log"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

var (
	httpClient *http.Client
)

// init HTTPClient
func init() {
	httpClient = createHTTPClient()
}

const (
	MaxIdleConns        int = 100
	MaxIdleConnsPerHost int = 100
	IdleConnTimeout     int = 90
)

// createHTTPClient for connection re-use
func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        MaxIdleConns,
			MaxIdleConnsPerHost: MaxIdleConnsPerHost,
			IdleConnTimeout:     time.Duration(IdleConnTimeout),
		},
	}
	return client
}

const (
	UrlPrefix = "http://172.21.33.7/"
)

func main() {
	// 定义接受参数的变量
	var id string
	var name string

	flag.StringVar(&id, "id", "360202198807090038", "用户id")
	flag.StringVar(&name, "name", "曾嵩", "用户name")

	// 解析命令行参数写入注册的flag里,在所有flag都注册之后，这一步一定不能少
	flag.Parse()

	var contentType = "application/x-www-form-urlencoded"
	// 养老保险
	api := UrlPrefix + "gateway/api/1/getdata"
	api = api + "?paperkind=110001&paperid=" + id + "&orgnum=0600&username=" + name
	fmt.Println("api::", api)

	basicInfo := httpDo(httpClient, "POST", api, contentType, "")
	log.Info("===>", basicInfo)

}

// 有时需要在请求的时候设置头参数、cookie之类的数据，就可以使用http.Do方法。
func httpDo(client *http.Client, requestType, url, contentType, data string) string {
	req, err := http.NewRequest(requestType, url, strings.NewReader(data))
	// req.Header.Set("Appkey", "718131542538321920")
	req.Header.Set("Appkey", "823877309051174912")
	if err != nil {
		return fmt.Sprintf("httpDo err:%v\n", err)
	}

	req.Header.Set("Content-Type", contentType)

	resp, e := client.Do(req)
	if e != nil {
		return fmt.Sprintf("接口调用出错：%v\n", e)
	}

	defer resp.Body.Close()

	body, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		return fmt.Sprintf("httpDo er:%v\n", er)
	}
	return string(body)
}
