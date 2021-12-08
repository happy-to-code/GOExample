package main

import (
	"bytes"
	"crypto/sha256"
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
	// contentType := "application/x-www-form-urlencoded"
	bxUrl := "http://localhost:8888/bx/test"

	appid := "TEST"                          // 应用接入代码
	timestamp := time.Now().UnixNano() / 1e6 // 13位时间戳
	signSecret := "abcd1234"                 // 签名秘钥

	identifynumber := "360202198807090038" // 客户证件号码
	insuredname := "曾嵩"                    // 客户名称

	signStr := fmt.Sprintf("%s%s%s%d", signSecret, identifynumber, insuredname, timestamp)
	// 	计算哈希（sha256）
	sum := sha256.Sum256([]byte(signStr))
	signature := fmt.Sprintf("%x", sum)
	fmt.Println("signature:", signature)

	var buffer bytes.Buffer
	buffer.WriteString(bxUrl)
	buffer.WriteString("&identifynumber=")
	buffer.WriteString(identifynumber)
	buffer.WriteString("&insuredname=")
	buffer.WriteString(insuredname)
	buffer.WriteString("&appid=")
	buffer.WriteString(appid)
	buffer.WriteString("&timestamp=")
	buffer.WriteString(fmt.Sprintf("%d", timestamp))
	buffer.WriteString("&signature=")
	buffer.WriteString(signature)
	bxUrl = buffer.String()
	fmt.Println("url:", bxUrl)

	headerSecurity := "appid=" + appid + ";timestamp=" + fmt.Sprintf("%d%s", timestamp, ";") + "signature=" + signature
	fmt.Println("headerSecurity:", headerSecurity)
}

// 有时需要在请求的时候设置头参数、cookie之类的数据，就可以使用http.Do方法。
func httpDo(client *http.Client, requestType, url, contentType, data string) string {
	req, err := http.NewRequest(requestType, url, strings.NewReader(data))
	// req.Header.Set("Appkey", "718131542538321920")
	// req.Header.Set("Appkey", "823877309051174912")

	// req.Header.Set("piccjs-svc-security", "823877309051174912")
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
