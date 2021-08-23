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
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Individual 个体信息
type Individual struct {
	AuthID      string `json:"authId" binding:"required"`
	SubjectType string `json:"subjectType" binding:"oneof=natural legal"`
	BankID      string `json:"bankId" binding:"required"`
	EndTime     string `json:"endTime" binding:"required"` // 格式为：yyyy-MM-dd hh:mm:ss
	FileHash    string `json:"fileHash" binding:"required"`
	// FilePath    string  `json:"filePath"`
	Subject Subject `json:"subject" binding:"required"`
}
type Subject struct {
	Id         string `json:"id" binding:"required"`   // 身份证号码or社会统一信用代码
	Name       string `json:"name" binding:"required"` // 姓名or企业/组织全称
	Code       string `json:"code"`                    // 组织机构代码
	WeChatName string `json:"weChatName"`
	Location   string `json:"location"`
	Ip         string `json:"ip"`
	Phone      string `json:"phone" binding:"required"`
}
type PersonInfo struct {
	Qlr  string `json:"qlr"`  // 权利人名字
	Sfzh string `json:"sfzh"` // 权利人身份证
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
