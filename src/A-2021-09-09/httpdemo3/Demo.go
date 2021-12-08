package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"reflect"
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
		Paperid:   "321027199004113325",
		Orgnum:    "0600",
		Username:  "宋飞宇",
	}
	var param2 = Param{
		Paperkind: "110001",
		Paperid:   "321027199004113325",
		Orgnum:    "0600",
		Username:  "王道萍",
	}
	var param3 = Param{
		Paperkind: "110001",
		Paperid:   "321088198610215450",
		Orgnum:    "0600",
		Username:  "吴正寅",
	}
	var param4 = Param{
		Paperkind: "110001",
		Paperid:   "321081198611200344",
		Orgnum:    "0600",
		Username:  "孙雅",
	}
	var param5 = Param{
		Paperkind: "110001",
		Paperid:   "360202198807090038",
		Orgnum:    "0600",
		Username:  "曾嵩",
	}
	var param6 = Param{
		Paperkind: "110001",
		Paperid:   "321088198811155757",
		Orgnum:    "0600",
		Username:  "周江安",
	}
	var param7 = Param{
		Paperkind: "110001",
		Paperid:   "321084198106080416",
		Orgnum:    "0600",
		Username:  "吕兵兵",
	}

	paramList = append(paramList, param1, param2, param3, param4, param5, param6, param7)

	var contentType = "application/x-www-form-urlencoded"
	// // 用户基本信息查询接口
	// api := UrlPrefix + "gateway/api/1/yhjbxxcxjk"
	//
	// // http://ip/gateway/api/1/yhjbxxcxjk
	// // ?bac045=&aae135&baz805&aac001=2000000309107370&date=20150917&time=142623
	// api = api + "?bac045=&aae135=" + param.Paperid + "&baz805&aac001=2000000309107370&date=20150917&time=142623"
	// fmt.Println("---->", api)
	// httpDo(httpClient, "POST", api, contentType, "")

	for _, param := range paramList {
		// 养老保险缴费记录查询接口
		// http://ip/gateway/api/1/ylbxjfjlcx
		// ?bac045=&aae135&baz805&aac001=2000000309107370&date=20150917&time=142623&cpage=1&rows=1
		api := UrlPrefix + "gateway/api/1/ylbxjfjlcx"
		api = api + "?bac045=&aae135=" + param.Paperid + "&baz805&aac001=&date=20210818&time=142623&cpage=0&rows=120"
		ylbx := httpDo(httpClient, "POST", api, contentType, "")
		fmt.Println("养老保险：", ylbx)

		// 养老保险
		res := new(PensionRes)
		json.Unmarshal([]byte(ylbx), res)
		if res.Code == 200 {
			pensionRes := handPensionRes(res.Data)
			bytes, _ := json.Marshal(pensionRes)
			fmt.Println("---->", string(bytes))

		}
		fmt.Println("----------------------------" + param.Username + "-------------------------------")
	}
}

// 有时需要在请求的时候设置头参数、cookie之类的数据，就可以使用http.Do方法。
func httpDo(client *http.Client, requestType, url, contentType, data string) string {
	req, err := http.NewRequest(requestType, url, strings.NewReader(data))
	// req.Header.Set("Appkey", "718131542538321920")
	req.Header.Set("Appkey", "823877309051174912")
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

// PensionRes 养老保险
type PensionRes struct {
	Code    int    `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}
type PensionRaw struct {
	Header struct {
		Date    string `json:"date"`
		SysCode string `json:"sysCode"`
		Ref     string `json:"ref"`
		Rst     struct {
			BusiCode  string `json:"busiCode"`
			TradeCode string `json:"tradeCode"`
			Info      string `json:"info"`
		} `json:"rst"`
		BusCode  string `json:"busCode"`
		ReSend   string `json:"reSend"`
		Sender   string `json:"sender"`
		TradeSrc string `json:"tradeSrc"`
		Time     string `json:"time"`
		Version  string `json:"version"`
		Reciver  string `json:"reciver"`
	} `json:"header"`
	Body struct {
		DataSet struct {
			RowSet struct {
				Row interface{} `json:"row"`
			} `json:"rowSet"`
			PageSize   string `json:"pageSize"`
			TotalCount string `json:"totalCount"`
			Lst        string `json:"lst"`
			PageNum    string `json:"pageNum"`
		} `json:"dataSet"`
	} `json:"body"`
}

type Pension struct {
	TotalMonth string          `json:"total_month"`
	Records    []PensionRecord `json:"records"`
}
type PensionRecord struct {
	Company       string `json:"company"`        // 缴纳单位名称
	PayMonth      string `json:"pay_month"`      // 缴纳月份
	PayDate       string `json:"pay_date"`       // 缴纳日期
	CompanyAmount string `json:"company_amount"` // 单位应缴金额
	PersonAmount  string `json:"person_amount"`  // 个人应缴金额
	TotalAmount   string `json:"total_amount"`   // 缴纳总额
	PersonBase    string `json:"person_base"`    // 人员缴费基数
	Status        string `json:"status"`         // 养老保险缴纳状态
}

func handPensionRes(data string) (pension Pension) {
	var pensionRaw PensionRaw
	err := json.Unmarshal([]byte(data), &pensionRaw)
	if err != nil {
		fmt.Printf("JSON反序列化出错:%v", err)
		return
	}
	if pensionRaw.Header.Rst.Info != "OK" {
		fmt.Println("未查询到任何信息")
		return
	}

	// 总共缴纳的月份
	pension.TotalMonth = pensionRaw.Body.DataSet.TotalCount
	// row原始数据
	row := pensionRaw.Body.DataSet.RowSet.Row
	rowBytes, _ := json.Marshal(row)

	var pensionRecords []PensionRecord
	rowTypeStr := fmt.Sprintf("%s", reflect.TypeOf(row))
	// 切片类型
	if rowTypeStr == "[]interface {}" {
		var rows []Row
		json.Unmarshal(rowBytes, &rows)
		for _, v := range rows {
			pensionRecord := PensionRecord{
				Company:       v.Aae044,
				PayMonth:      v.Aae003,
				PayDate:       v.Aae079,
				CompanyAmount: v.Aae080,
				PersonAmount:  v.Aae082,
				TotalAmount:   v.Aae019,
				PersonBase:    v.Aae180,
				Status:        v.Aaa115,
			}
			pensionRecords = append(pensionRecords, pensionRecord)
		}
	} else if rowTypeStr == "map[string]interface {}" {
		var v Row
		json.Unmarshal(rowBytes, &v)
		pensionRecord := PensionRecord{
			Company:       v.Aae044,
			PayMonth:      v.Aae003,
			PayDate:       v.Aae079,
			CompanyAmount: v.Aae080,
			PersonAmount:  v.Aae082,
			TotalAmount:   v.Aae019,
			PersonBase:    v.Aae180,
			Status:        v.Aaa115,
		}
		pensionRecords = append(pensionRecords, pensionRecord)
	} else {
		return
	}

	pension.Records = pensionRecords
	return
}

type Row struct {
	Aae044 string `json:"aae044"` // 单位名称
	Bab001 string `json:"bab001"` //
	Aae140 string `json:"aae140"` //
	Aae003 string `json:"aae003"` // 最近缴纳月份
	Aae079 string `json:"aae079"` // 到账年月日
	Aae078 string `json:"aae078"` // 缴纳状态，未缴纳为0，已缴纳为1
	Aae080 string `json:"aae080"` // 单位应缴金额
	Aae019 string `json:"aae019"` // 缴纳总额
	Aaa115 string `json:"aaa115"` // 养老保险缴纳状态？
	Aae082 string `json:"aae082"` // 个人应缴金额
	Aae180 string `json:"aae180"` // 人员缴费基数
}
