package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
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

type PostDemo struct {
	Name string `1.json:"name"`
	Age  int    `1.json:"age"`
}

// Individual 个体信息
type Individual struct {
	AuthID      string `1.json:"authId" binding:"required"`
	SubjectType string `1.json:"subjectType" binding:"oneof=natural legal"`
	BankID      string `1.json:"bankId" binding:"required"`
	EndTime     string `1.json:"endTime" binding:"required"` // 格式为：yyyy-MM-dd hh:mm:ss
	FileHash    string `1.json:"fileHash" binding:"required"`
	// FilePath    string  `1.json:"filePath"`
	Subject Subject `1.json:"subject" binding:"required"`
}
type Subject struct {
	Id         string `1.json:"id" binding:"required"`   // 身份证号码or社会统一信用代码
	Name       string `1.json:"name" binding:"required"` // 姓名or企业/组织全称
	Code       string `1.json:"code"`                    // 组织机构代码
	WeChatName string `1.json:"weChatName"`
	Location   string `1.json:"location"`
	Ip         string `1.json:"ip"`
	Phone      string `1.json:"phone" binding:"required"`
}
type PersonInfo struct {
	Qlr  string `1.json:"qlr"`  // 权利人名字
	Sfzh string `1.json:"sfzh"` // 权利人身份证
}

const (
	UrlPrefix = "http://172.21.47.49/"
)

func main() {
	var gsList = []string{"仪征市都市壹壹捌商务宾馆"}
	for _, gs := range gsList {
		gs = url.QueryEscape(gs)

		// 社会法人社保缴费信息接口
		api := UrlPrefix + "gateway/api/1/shfrsbjfxxcxjk"

		api = api + "?jgqc=" + gs
		fmt.Println("url:", api)
		httpDo(httpClient, "GET", api, "application/x-www-form-urlencoded", "")

		// 社会法人社保欠缴信息查询接口
		api = UrlPrefix + "gateway/api/1/shfrsbqjxxcxjk"
		api = api + "?jgqc=" + gs
		fmt.Println("url:", api)
		httpDo(httpClient, "GET", api, "application/x-www-form-urlencoded", "")
	}
}

// 有时需要在请求的时候设置头参数、cookie之类的数据，就可以使用http.Do方法。
func httpDo(client *http.Client, requestType, url, contentType, data string) {
	req, err := http.NewRequest(requestType, url, strings.NewReader(data))
	req.Header.Set("Appkey", "718131542538321920")
	// req.Header.Set("Appkey", "823877309051174912")
	if err != nil {
		fmt.Printf("httpDo err:%v\n", err)
	}

	// fmt.Printf("req:888888888888*********************====>%v\n", req)
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Set("Cookie", "name=anny")
	// req.Header.Set("Content-Type", contentType)

	resp, e := client.Do(req)
	if e != nil {
		fmt.Printf("接口调用出错：%v\n", e)
	}

	defer resp.Body.Close()

	body, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		fmt.Printf("httpDo er:%v\n", er)
	}
	fmt.Println("====>httpDo返回数据：", string(body))
}
