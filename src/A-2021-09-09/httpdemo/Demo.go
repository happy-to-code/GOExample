package main

import (
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

	// var param1 = Param{
	// 	Paperkind: "110001",
	// 	Paperid:   "321027199004113325",
	// 	Orgnum:    "0600",
	// 	Username:  "宋飞宇",
	// }
	// var param2 = Param{
	// 	Paperkind: "110001",
	// 	Paperid:   "321027199004113325",
	// 	Orgnum:    "0600",
	// 	Username:  "王道萍",
	// }
	// var param3 = Param{
	// 	Paperkind: "110001",
	// 	Paperid:   "321088198610215450",
	// 	Orgnum:    "0600",
	// 	Username:  "吴正寅",
	// }
	// var param4 = Param{
	// 	Paperkind: "110001",
	// 	Paperid:   "321081198611200344",
	// 	Orgnum:    "0600",
	// 	Username:  "孙雅",
	// }
	var param5 = Param{
		Paperkind: "110001",
		Paperid:   "360202198807090038",
		Orgnum:    "0600",
		Username:  "曾嵩",
	}
	// var param6 = Param{
	// 	Paperkind: "110001",
	// 	Paperid:   "321088198811155757",
	// 	Orgnum:    "0600",
	// 	Username:  "周江安",
	// }
	// var param7 = Param{
	// 	Paperkind: "110001",
	// 	Paperid:   "321084198106080416",
	// 	Orgnum:    "0600",
	// 	Username:  "吕兵兵",
	// }

	// paramList = append(paramList, param1, param2, param3, param4, param5, param6, param7)
	paramList = append(paramList, param5)

	for i, param := range paramList {
		api := UrlPrefix + "gateway/api/1/getdata"

		// pBytes, _ := json.Marshal(param)
		// var contentType = "application/x-www-form-urlencoded"
		var contentType = "application/x-www-form-urlencoded"

		api = api + "?paperkind=" + param.Paperkind + "&paperid=" + param.Paperid + "&orgnum=" + param.Orgnum + "&username=" + param.Username
		fmt.Println("---->", api)
		httpDo(httpClient, "POST", api, contentType, "")
		fmt.Println("---------------------------fg----------------------------------", i)
	}

}

// 有时需要在请求的时候设置头参数、cookie之类的数据，就可以使用http.Do方法。
func httpDo(client *http.Client, requestType, url, contentType, data string) {
	req, err := http.NewRequest(requestType, url, strings.NewReader(data))
	// req.Header.Set("Appkey", "718131542538321920")
	req.Header.Set("Appkey", "823877309051174912")
	if err != nil {
		fmt.Printf("httpDo err:%v\n", err)
	}

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
