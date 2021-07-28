package main

import (
	"encoding/json"
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
	var inds []Individual

	var individual Individual
	var subject Subject
	subject.Id = "360202198807090038"
	subject.Name = "曾嵩"
	// subject.Id = "321088198811155757"
	// subject.Name = "周江安"
	individual.SubjectType = "natural"
	individual.Subject = subject

	// var individual2 Individual
	// var subject2 Subject
	// subject2.Id = "321088198610215450"
	// subject2.Name = "吴正寅"
	// individual2.SubjectType = "natural"
	// individual2.Subject = subject2

	// inds = append(inds, individual2)
	inds = append(inds, individual)

	for _, i := range inds {
		// requestType := "POST"
		contentType := "application/x-www-form-urlencoded"
		var p = PersonInfo{Qlr: i.Subject.Name, Sfzh: i.Subject.Id}
		bytes, _ := json.Marshal(p)
		parm := "args={\"selarea\":\"\",\"clientusername\":\"1\",\"clientusercid\":\"2\",\"ytcn\":\"test\",\"computerid\":\"4\",\"computermac\":\"5\",\"computername\":\"dsj\",\"psw\":\"7\",\"cxrzp\":\"8\",\"qlrlist\":[" + string(bytes) + "]}"
		// fmt.Println("参数 parm:", parm)

		// 覆盖范围邗江区、广陵区、经济开发区和江都区
		var api = UrlPrefix + "gateway/api/1/fwqscx"
		httpDo(httpClient, "POST", api, contentType, parm)

		// 婚姻
		var api1 = UrlPrefix + "gateway/api/1/hydjxxcx"
		api1 = api1 + "?gmsfzh=" + i.Subject.Id
		httpDo(httpClient, "GET", api1, "application/x-www-form-urlencoded", "")

		// 社保
		var api2 = UrlPrefix + "gateway/api/1/zrrsbjnxxcxjk"
		data := "xm=" + i.Subject.Name + "&zjhm=" + i.Subject.Id
		httpDo(httpClient, "POST", api2, "application/x-www-form-urlencoded", data)

		// 自然人严重失信黑名单
		// var api = UrlPrefix + "gateway/api/1/zrryzsxhmdfy"
		// api = api + "?ROWNUM=0&PAGESIZE=1&SFZJHM=" + i.Subject.Id
		// httpDo(httpClient, "GET", api, "application/x-www-form-urlencoded", "")
		//
		// // 自然人守信红名单Url
		// api = UrlPrefix + "gateway/api/1/shfrcssxhmdfyfw"
		// // api = api + "?ROWNUM=0&PAGESIZE=1&QYMC=" + i.Subject.Name
		// data := "ROWNUM=0&PAGESIZE=1&QYMC=" + i.Subject.Name
		// httpDo(httpClient, "GET", api, "application/x-www-form-urlencoded", data)

		// 	网格数据
		fmt.Println("i==>", i)
		// api := "http://127.0.0.1:60005/fast-api/peasantHousehold/getDataByCardNo?cardNo=321022195706281518"
		// api := "http://172.21.43.41:60005/fast-api/peasantHousehold/getDataByCardNo?cardNo=321022195706281518"
		var api4 = "http://172.21.43.41:60005/fast-api/peasantHousehold/getDataByCardNo?cardNo=321022197206221518"
		httpDo(httpClient, "GET", api4, "application/x-www-form-urlencoded", "")

		// api := "http://127.0.0.1:8088/fast-api/peasantHousehold/getDataList?beginTime="+url.QueryEscape("2021-06-13 00:00:00")+"&endTime="+url.QueryEscape("2021-06-23 00:00:00")
		// // data := "beginTime=2021-06-13 00:00:00&endTime=2021-06-23 00:00:00"
		// httpDo(httpClient, "GET", api, "application/json;charset=UTF-8", "")

	}

}

// 有时需要在请求的时候设置头参数、cookie之类的数据，就可以使用http.Do方法。
func httpDo(client *http.Client, requestType, url, contentType, data string) {
	req, err := http.NewRequest(requestType, url, strings.NewReader(data))
	req.Header.Set("Appkey", "718131542538321920")
	if err != nil {
		fmt.Printf("httpDo err:%v\n", err)
	}

	fmt.Printf("req:888888888888*********************====>%v\n", req)
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Set("Cookie", "name=anny")
	req.Header.Set("Content-Type", contentType)

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

func httpGet(url string) {
	header := http.Header{}
	header.Set("Appkey", "718131542538321920")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("httpGet err:%v\n", err)
	}

	defer resp.Body.Close()
	body, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		fmt.Printf("httpGet er:%v\n", er)
	}

	fmt.Println(string(body))
}

// 使用这个方法的话，第二个参数要设置成”application/x-www-form-urlencoded”
// 否则post参数无法传递。
func httpPost(url string) {
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Printf("httpPost err:%v\n", err)
	}

	defer resp.Body.Close()
	body, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		fmt.Printf("httpPost er:%v\n", er)
	}

	fmt.Println(string(body))
}

func httpPostForm(u, k, v string) {
	header := http.Header{}
	header.Set("Appkey", "718131542538321920")
	resp, err := http.PostForm(u, url.Values{k: {v}, "Appkey": {"718131542538321920"}})

	if err != nil {
		fmt.Printf("httpPostForm err:%v\n", err)
	}

	defer resp.Body.Close()
	body, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		fmt.Printf("httpPostForm er:%v\n", er)
	}

	fmt.Println("httpPostForm==>", string(body))
}
