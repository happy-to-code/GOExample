package main

import (
	"encoding/json"
	"fmt"
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

type Param struct {
	Paperkind string `json:"paperkind"`
	Paperid   string `json:"paperid"`
	Orgnum    string `json:"orgnum"`
	Username  string `json:"username"`
}

func main() {
	// gs = url.QueryEscape(gs)

	var paramList []Param

	var param1 = Param{
		Paperkind: "110001",
		Paperid:   "321084196601071745",
		Orgnum:    "0600",
		Username:  "李者女",
	}
	var param2 = Param{
		Paperkind: "110001",
		Paperid:   "32102219660314171X",
		Orgnum:    "0600",
		Username:  "陆明吾",
	}

	contentType := "application/x-www-form-urlencoded"
	paramList = append(paramList, param1, param2)
	for _, param := range paramList {
		// 3区
		t0 := time.Now().Unix()
		str := queryFW(param, contentType, "args=", "gateway/api/2/fwqscx")
		t1 := time.Now().Unix()
		fmt.Println("3区：", str)
		fmt.Println("3区：", t1-t0)

		// 高邮
		t2 := time.Now().Unix()
		str2 := queryFW(param, contentType, "raw=", "gateway/api/1/gysfwqscxjk")
		t3 := time.Now().Unix()
		fmt.Println("高邮：", str2)
		fmt.Println("高邮：", t3-t2)

		// 仪征
		t4 := time.Now().Unix()
		str3 := queryFW(param, contentType, "raw=", "gateway/api/1/yzsfwqscxjk")
		t5 := time.Now().Unix()
		fmt.Println("仪征：", str3)
		fmt.Println("仪征：", t5-t4)

		fmt.Println("---------------------------------------------")
	}
}

func queryFW(param Param, contentType string, arg string, url string) string {
	type stru struct {
		Qlr  string `json:"qlr"`
		Sfzh string `json:"sfzh"`
	}
	p := stru{
		Qlr:  param.Username,
		Sfzh: param.Paperid,
	}

	bytes, _ := json.Marshal(p)
	pa := arg + "{\"selarea\":\"\",\"clientusername\":\"1\",\"clientusercid\":\"2\",\"ytcn\":\"test\",\"computerid\":\"4\",\"computermac\":\"5\",\"computername\":\"dsj\",\"psw\":\"7\",\"cxrzp\":\"8\",\"qlrlist\":[" + string(bytes) + "]}"

	// parm := fmt.Sprintf(param, string(bytes))
	var api = "http://172.21.47.49/" + url
	// ylbx := httpDo(httpClient, "POST", api, contentType, "")
	str := httpDo(httpClient, "POST", api, contentType, pa)
	return str
}

// 有时需要在请求的时候设置头参数、cookie之类的数据，就可以使用http.Do方法。
func httpDo(client *http.Client, requestType, url, contentType, data string) string {
	req, err := http.NewRequest(requestType, url, strings.NewReader(data))
	req.Header.Set("Appkey", "718131542538321920")
	// req.Header.Set("Appkey", "823877309051174912")
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
