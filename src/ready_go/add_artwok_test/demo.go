package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"
)

var param = `{
  "artistNo": "no12345678",
  "name": "小狗旺财",
  "artworkNo": "No110_",
  "author": "胡朝水",
  "size": "1360cm*68cm",
  "creationTime": "1200",
  "category": "书画类",
  "type": "风景",
  "material": "宣纸",
  "fullArtwork": {
    "hash": "5543749626770f0e459055cca66a2bca38815ef94cbc34dae9c86bf112768555",
    "address": "http://10.1.3.5/url/addr11",
    "name": "艺术品全图"
  },
  "signature": {
    "hash": "8883749626770f0e459067cca66a2bca45815ef94cbc34dae9c86bf112768696",
    "address": "http://10.1.3.5/url/addr22",
    "name": "艺术品落款"
  },
  "seal": {
    "hash": "5763749626770f0e459067cca66a2bca99915ef94cbc34dae9c86bf112768eee",
    "address": "http://10.1.3.5/url/addr33",
    "name": "艺术品名章"
  },
  "groupPhoto": {
    "hash": "kjk3749626770f0e459067cca66a2b8838815ef94cbc34dae9c86bf112768fmm",
    "address": "http://10.1.3.5/url/addr44",
    "name": "艺术品和艺术家合照"
  },
  "featurePointBitmap": [
    {
      "hash": "ff43749626770f0e459067cca66a2bca38815ef94cbc34dae9c86bf112768bbb",
      "address": "http://10.1.3.5/url/point1",
      "coordinate": [
        "35.89",
        "90.11"
      ],
      "type": "类型1"
    },
    {
      "hash": "eee3749626770f0e4590tttca66a2bca38815ef94cbc34dae9c86bf112768xtb",
      "address": "http://10.1.3.5/url/point2",
      "coordinate": [
        "85.89",
        "80.41"
      ],
      "type": "类型2"
    }
  ],
  "registrationCertificate": {
    "hash": "fffff49626770f0e459067cca76a2bca38815ef941bc34dae9c86bf112768f88",
    "address": "http://10.1.3.5/url/addr55",
    "name": "作品登记证书"
  },
  "otherAuthCertificate": [
    {
      "hash": "8943749626770f0e459067cca66a2bcb38815ef94cbc34dae9c86bf112768bbc",
      "address": "http://10.1.3.5/url/addr66",
      "name": "其他认证证书"
    }
  ],
  "squareFeet": 10,
  "style": "写实",
  "theme": "风景",
  "formOfExpression": "水墨",
  "tradeInfoList": [
    {
      "artworkNo": "no458990",
      "transactionDate": 1551804856,
      "transactionAmount": "80000万",
      "seller": "张三",
      "buyer": "李*",
      "assuranceUnit": "弘艺丰",
      "depositTime": 1551804856,
      "orderNo": "no4590000"
    },
    {
      "artworkNo": "no458988",
      "transactionDate": 1551004856,
      "transactionAmount": "82000万",
      "seller": "王五",
      "buyer": "钱**",
      "assuranceUnit": "弘艺丰",
      "depositTime": 1551004856,
      "orderNo": "no4596609"
    }
  ],
  "agency": {
    "name": "苏州书画协会",
    "licenseRegistrationNo": "110108001372301",
    "organizationCode": "79340139-9",
    "address": "北京市东城区",
    "businessTerm": 15,
    "applicant": "张三",
    "jobTitle": "画家",
    "documentType": "身份证",
    "documentNo": "371***************",
    "phone": "1381254****",
    "email": "zhangsan@163.com",
    "executiveStandard": "T/CSIQ1011-2015-M",
    "certificateType": "备案证书"
  }
}`

func main() {
	var url = "http://10.1.3.151:9100/hyf/v1/artwork/add"

	var mParm map[string]interface{}
	err := json.Unmarshal([]byte(param), &mParm)
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		intn := rand.Intn(10000000)
		randstr := fmt.Sprintf("%d", intn)
		mParm["name"] = "小狗旺财_" + randstr
		mParm["artworkNo"] = "No110_" + randstr

		bytes, _ := json.Marshal(mParm)

		err, resp := httpDo(createHTTPClient(), "POST", url, "application/json; charset=utf-8", string(bytes))
		if err != nil {
			fmt.Println("add failed:", err)
			continue
		}
		fmt.Printf("[%d][%s]resp:%s\n", i+1, randstr, resp)
	}

	// getUrl := "http://10.1.3.151:9100/v2/tx/detail/6c7dd1af885a6b9b35ad3a8814cdf58e8fd4a1b468b9136a0445b98efc26d56e?ledger=yrn3h7x1cf"
	// err, resp := httpDo(createHTTPClient(), "GET", getUrl, "application/json", string(""))
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(resp)

}

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
