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
	UrlPrefix = "http://localhost:8899/testParameter"
)

func main() {

	var contentType = "application/json"
	m := make(map[string]string)
	m["id"] = "123"
	m["didJsonStr"] = `{
    "@context": "https://w3id.org/did/v1",
    "id": "did:example:3mYESp9TnqVfkTqPpQb5jAEkWsu8",
    "created": "2021-12-08T10:17:08Z",
    "updated": "2021-12-08T10:17:08Z",
    "publicKey": [
        {
            "id": "did:example:3mYESp9TnqVfkTqPpQb5jAEkWsu8#key-1",
            "type": "EC",
            "publicKeyHex": "3059301306072a8648ce3d020106082a811ccf5501822d03420004c7b59b0b845ccb7cfacbe2b867fcd40f370828c2a2da0835aa8b693a557b34ba2b3e5cf489ede26d9cf9b63155e8d6c3ff57ffe5a371e8340f93bf7a232b46d1"
        }
    ],
    "authentication": [
        "did:example:3mYESp9TnqVfkTqPpQb5jAEkWsu8#key-1"
    ],
    "proof": {
        "type": "SM3WITHSM2",
        "creator": "did:example:3mYESp9TnqVfkTqPpQb5jAEkWsu8#key-1",
        "signatureValue": "MEQCIHnx6n+ykHY65MBcadi5rkV613WXj2wCatnYVqjxYL7EAiBaPDc+ZZ4HFoAgRA6dscJFCc829CUXgyCAal2OV14kww=="
    }
}`

	bytes, _ := json.Marshal(m)

	httpDo(httpClient, "POST", UrlPrefix, contentType, string(bytes))

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
