package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

var httpClient *http.Client

func init() {
	httpClient = createHTTPClient()
}

type TestStruct struct {
	Id        int   `json:"id,omitempty"`
	Timestamp int64 `json:"timestamp,omitempty"`
}

type TestData struct {
	Data string `json:"data,omitempty"`
}

type ChainResponse struct {
	State   int    `json:"state"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func main() {
	var wg sync.WaitGroup

	url := "http://36.137.215.142:48180/"
	storeUrl := url + "v2/tx/store?sync=true"
	getUrl := url + "v2/tx/detail/"
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go testSyncStore(i, storeUrl, getUrl, &wg)
	}
	wg.Wait()
}

func testSyncStore(i int, storeUrl string, getUrl string, wg *sync.WaitGroup) {
	defer wg.Done()
	var t = TestStruct{Id: i, Timestamp: time.Now().UnixNano()}
	bytes, _ := json.Marshal(t)
	var data = TestData{Data: string(bytes)}
	dataBytes, _ := json.Marshal(data)

	err, s := httpDo(httpClient, "POST", storeUrl, "application/json", string(dataBytes))
	if err != nil {
		panic(err)
	}
	// fmt.Println("s1====>", s)
	// {"state":200,"message":"success","data":"f94a084701de49d3e2a3f139923a927de926f82e08c0fdd3b3bc0e81ae6182b6"}
	var resp ChainResponse
	err = json.Unmarshal([]byte(s), &resp)
	if err != nil {
		fmt.Printf("[%s]反序列化失败:%v\n", s, err)
		return
	}
	txid := resp.Data
	getUrl2 := getUrl + txid
	// fmt.Println("getUrl::::::", getUrl2)
	err, s2 := httpDo(httpClient, "GET", getUrl2, "application/json", "")
	if err != nil {
		fmt.Printf("[%s]获取交易详情失败，%v\n", txid, err)
		return
	}
	fmt.Println("s2:::::::", s2)
	return
}

// MySQL数据库配置
const (
	MaxIdleConns        int = 100
	MaxIdleConnsPerHost int = 100
	IdleConnTimeout     int = 90
)

// createHTTPClient for connection re-use
func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				// Timeout:   5 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        MaxIdleConns,
			MaxIdleConnsPerHost: MaxIdleConnsPerHost,
			IdleConnTimeout:     time.Duration(IdleConnTimeout),
		},
		Timeout: time.Duration(30) * time.Second,
	}
	return client
}

func httpDo(client *http.Client, requestType, url, contentType, data string) (error, string) {
	req, err := http.NewRequest(requestType, url, strings.NewReader(data))
	if err != nil {
		return fmt.Errorf("httpDo err:%v\n", err), ""
	}

	req.Header.Set("Content-Type", contentType)

	resp, e := client.Do(req)
	if e != nil {
		return fmt.Errorf("接口调用出错：%v\n", e), ""
	}

	defer resp.Body.Close()

	body, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		return fmt.Errorf("httpDo er:%v\n", er), ""
	}
	return nil, string(body)
}
