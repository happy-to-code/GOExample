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

const (
	MaxIdleConns        int = 100
	MaxIdleConnsPerHost int = 100
	IdleConnTimeout     int = 90
)

var (
	httpClient *http.Client
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

// 有时需要在请求的时候设置头参数、cookie之类的数据，就可以使用http.Do方法。
func httpDo(client *http.Client, ty, url, data string) (string, error) {
	req, err := http.NewRequest(ty, url, strings.NewReader(data))
	if err != nil {
		return "", fmt.Errorf("NewRequest err:%v\n", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, e := client.Do(req)
	if e != nil {
		return "", fmt.Errorf("接口调用出错：%v\n", e)
	}

	defer resp.Body.Close()

	body, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		return "", fmt.Errorf("ioutil.ReadAll err:%v\n", er)
	}

	// fmt.Println("1===============================1===========================1")
	// var reply Reply
	// 1.json.Unmarshal(body, &reply)
	// log.Printf("reply:%+v\n", reply)
	// fmt.Println("2===============================2===========================2")

	return string(body), nil
}

func init() {
	// 初始化httpClient
	httpClient = createHTTPClient()
}

func main() {

	var s = `
{
    "SubjectType": "natural",
    "EndTime": "2021-11-11 00:00:18",
    "FileHash": "wetest",
    "Subject": {
        "id": "360202198807090038",
        "name": "曾嵩"
    }
}
`
	t1 := time.Now().Unix()
	fmt.Println(t1)
	var add = "http://10.1.3.150:9100/jml/v1/authorize/add?sync=true"
	do, err := httpDo(httpClient, "POST", add, s)
	if err != nil {
		panic(err)
	}
	t2 := time.Now().Unix()
	fmt.Println(t2)
	var res Response
	json.Unmarshal([]byte(do), &res)
	txId := res.Data.TxId
	fmt.Println(txId, t2-t1)

	// var detail = "http://10.1.3.150:9100/v2/tx/detail/6631f7f85d674d607218bdd51d07fd4ce5ee8e2a3802c7cbe9efa10f9bc81f80?ledger=et0cgum17f"
	var detail = "http://10.1.3.150:9100/v2/tx/detail/" + txId + "?ledger=et0cgum17f"
	fmt.Println(detail)
	// time.Sleep(time.Second*3)
	s2, err := httpDo(httpClient, "GET", detail, "")
	if err != nil {
		panic(err)
	}
	fmt.Println("s2", s2)

}

type Response struct {
	State   int    `1.json:"state,omitempty"`
	Message string `1.json:"message,omitempty"`
	Data    Data   `1.json:"data"`
}

type Data struct {
	TxId   string `1.json:"txId,omitempty"`
	AuthId string `1.json:"authId,omitempty"`
}
