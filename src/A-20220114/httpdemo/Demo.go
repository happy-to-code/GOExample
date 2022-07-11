package main

import (
	"encoding/json"
	"fmt"
	"github.com/siddontang/go/log"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
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
	UrlPrefix = "http://10.1.3.116:3334"
)

func main() {
	var k = 464634748
	// for i := 40000; i < 45000; i++ {
	for true {
		log.Info("kkkkkk:", k)
		blockNumStr := strconv.FormatInt(int64(k), 16)

		var data = `{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x%s", true],"id":1}`
		data = fmt.Sprintf(data, blockNumStr)
		res := httpDo(httpClient, "POST", UrlPrefix, "application/1.json", data)

		var blockObject BlockObject
		err := json.Unmarshal([]byte(res), &blockObject)
		if err != nil {
			log.Warn(res, ",反序列化出错：", err)
		} else if blockObject.Jsonrpc != "" && len(blockObject.Result.Transactions) > 0 {
			// log.Infof("k:%d,blockObject==>%+v", k, blockObject)
			// return
			for _, transaction := range blockObject.Result.Transactions {
				if transaction.To == "0x" {
					log.Info("k:", k)
					log.Info("res==>", res)
					return
				}
			}
		}
		k--
	}

}

// 有时需要在请求的时候设置头参数、cookie之类的数据，就可以使用http.Do方法。
func httpDo(client *http.Client, requestType, url, contentType, data string) string {
	req, err := http.NewRequest(requestType, url, strings.NewReader(data))
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

type BlockObject struct {
	Jsonrpc string `1.json:"jsonrpc"`
	Id      int    `1.json:"id"`
	Result  Result `1.json:"result"`
}

type Result struct {
	Difficulty       string        `1.json:"difficulty"`
	ExtraData        string        `1.json:"extraData"`
	GasLimit         string        `1.json:"gasLimit"`
	GasUsed          string        `1.json:"gasUsed"`
	Hash             string        `1.json:"hash"`
	LogsBloom        string        `1.json:"logsBloom"`
	Miner            string        `1.json:"miner"`
	MixHash          string        `1.json:"mixHash"`
	Nonce            string        `1.json:"nonce"`
	Number           string        `1.json:"number"`
	ParentHash       string        `1.json:"parentHash"`
	ReceiptsRoot     string        `1.json:"receiptsRoot"`
	Sha3Uncles       string        `1.json:"sha3Uncles"`
	Size             string        `1.json:"size"`
	StateRoot        string        `1.json:"stateRoot"`
	Timestamp        string        `1.json:"timestamp"`
	TotalDifficulty  string        `1.json:"totalDifficulty"`
	Transactions     []Transaction `1.json:"transactions"`
	TransactionsRoot string        `1.json:"transactionsRoot"`
	Uncles           []string      `1.json:"uncles"`
}

type Transaction struct {
	BlockHash        string `1.json:"blockHash"`
	BlockNumber      string `1.json:"blockNumber"`
	From             string `1.json:"from"`
	Gas              string `1.json:"gas"`
	GasPrice         string `1.json:"gasPrice"`
	Hash             string `1.json:"hash"`
	Input            string `1.json:"input"`
	Nonce            string `1.json:"nonce"`
	To               string `1.json:"to"`
	TransactionIndex string `1.json:"transactionIndex"`
	Value            string `1.json:"value"`
	Type             string `1.json:"type"`
	V                string `1.json:"v"`
	R                string `1.json:"r"`
	S                string `1.json:"s"`
}
