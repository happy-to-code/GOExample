package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/martian/log"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"mime/multipart"
	"time"
)

type PostDemo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// fmt.Println("*********************************")
	// // http://localhost:9088/v1/api/world
	// // url := "http://localhost:9088/v1/api/world"
	// url := "http://172.21.47.49:80/gateway/api/1/mysxjk"
	//
	// // err, s := httpGet(url)
	// s, err := httpGetAndPost(url, "POST", "")
	// if err == nil {
	// 	// panic(err)
	// 	fmt.Println(s)
	// } else {
	// 	fmt.Println("err:", err)
	// }
	//
	// fmt.Println("1--------------------------")
	// url = "http://172.21.47.49:80/gateway/api/1/hydjxxcx?gmsfzh=360202198807090038"
	// s3, err3 := httpGetAndPost(url, "POST", "")
	// if err3 == nil {
	// 	// panic(err)
	// 	fmt.Println(s3)
	// } else {
	// 	fmt.Println("err:", err3)
	// }
	// fmt.Println("2--------------------------")
	//
	// // ========================================================
	// // url := "http://localhost:9088/v1/api/post"
	// // url := "http://localhost:9088/v1/api/post"
	// // for i := 0; i < 10; i++ {
	// // 	var p PostDemo
	// // 	p.Age = i
	// // 	p.Name = "aa" + strconv.Itoa(i)
	// // 	jsonBytes, _ := json.Marshal(p)
	// // 	// fmt.Println("---->", string(jsonBytes))
	// // 	// err, s := httpPost(url, string(jsonBytes))
	// // 	s, err := httpGetAndPost(url,"POST",string(jsonBytes))
	// // 	if err == nil {
	// // 		fmt.Println(s)
	// // 	}
	// // }

	name2url := make(map[string]string)
	// name2url["房屋权属信息"] = "1"
	// name2url["社会法人社保缴费信息"] = "2"
	// name2url["社会法人社保欠缴信息"] = "3"
	name2url["自然人社保缴纳信息"] = "4"
	// name2url["婚姻登记信息"] = "5"
	// name2url["自然人守信红名单"] = "6"
	// name2url["社会法人守信红名单"] = "7"
	// name2url["自然人严重失信黑名单"] = "8"

	var individual Individual
	var subject Subject
	subject.Id = "360202198807090038"
	subject.Name = "曾嵩"

	// subject.Id = "321088198610215450"
	// subject.Name = "吴正寅"

	// subject.Id = "321087198010127020"
	// subject.Name = "秦志梅"

	// subject.Id = "321081198611200344"
	// subject.Name = "孙雅"

	individual.SubjectType = "natural"
	individual.Subject = subject
	GetInfoFromBigData(name2url, individual)

}

func httpPost(url, jsonData string) (error, string) {
	req := &fasthttp.Request{}
	req.Header.Set("Appkey", "718131542538321920")
	req.SetRequestURI(url)

	// requestBody := []byte(jsonData)
	// req.SetBody(requestBody)

	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	// req.Header.SetContentType("application/x-www-form-urlencoded")
	req.Header.SetMethod("POST")

	resp := &fasthttp.Response{}

	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		fmt.Println("请求失败：", err)
		return err, ""
	}

	b := resp.Body()
	fmt.Println("result:\r\n", string(b))
	return nil, string(b)
}

func httpGet(url string) (error, string) {
	status, resp, err := fasthttp.Get(nil, url)
	if err != nil {
		fmt.Println("请求失败:", err.Error())
		return err, ""
	}

	if status != fasthttp.StatusOK {
		fmt.Println("请求没有成功:", status)
		return fmt.Errorf("%s,status为:%d", "请求没有成功", status), ""
	}

	// fmt.Println(string(resp))
	return nil, string(resp)
}

func httpGetAndPost(url, method, data, contentType string, timeOut int) (string, error) {
	if timeOut <= 0 {
		timeOut = 10
	}

	req := &fasthttp.Request{} // 相当于获取一个对象

	req.SetRequestURI(url) // 设置请求的url
	req.Header.Set("Appkey", "718131542538321920")

	// bytes, err := json.Marshal(data) // data是请求数据
	//
	// if err != nil {
	// 	return "", err
	// }

	req.SetBody([]byte(data)) // 存储转换好的数据

	// req.Header.SetContentType("application/json") // 设置header头信息
	req.Header.SetContentType(contentType) // 设置header头信息

	req.Header.SetMethod(method) // 设置请求方法

	fmt.Println("request================>:", req)

	resp := &fasthttp.Response{} // 相应结果的对象

	client := &fasthttp.Client{} // 发起请求的对象

	// if err := client.Do(req, resp); err != nil {
	// 	return "", err
	// }
	if err := client.DoTimeout(req, resp, time.Duration(timeOut)*time.Second); err != nil {
		return "", err
	}
	fmt.Println("resp.String()----->", resp.String())
	fmt.Println("resp.StatusCode()----->", resp.StatusCode())
	return string(resp.Body()), nil
}

const (
	UrlPrefix = "http://172.21.47.49/"
)

// GetInfoFromBigData 调用大数据局接口获取数据
func GetInfoFromBigData(name2url map[string]string, individual Individual) {
	// 返回对象
	// 接口调用超时时间
	var timeOut = viper.GetInt("http.timeout")
	// var contentType = "application/x-www-form-urlencoded"
	// var contentType = "multipart/form-data"
	// var contentType = "application/json"
	var p = PersonInfo{Qlr: individual.Subject.Name, Sfzh: individual.Subject.Id}
	pBytes, _ := json.Marshal(p)

	s := "{\"selarea\":\"\",\"clientusername\":\"1\",\"clientusercid\":\"2\",\"ytcn\":\"test\",\"computerid\":\"4\",\"computermac\":\"5\",\"computername\":\"dsj\",\"psw\":\"7\",\"cxrzp\":\"8\",\"qlrlist\":[" + string(pBytes) + "]}"

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	w.WriteField("args", s)
	//
	// // xm=" + individual.Subject.Name + "&zjhm=" + individual.Subject.Id
	//
	// w.WriteField("xm", individual.Subject.Name)
	// w.WriteField("zjhm", individual.Subject.Id)
	// //
	// w.Close()
	var contentType = w.FormDataContentType()
	// fmt.Println("bbbbb---->", b.String())

	for name, url := range name2url {
		log.Infof("name:%s,url:%s", name, url)
		m := appendUrl(name, url, individual)
		for api, requestMode := range m {
			if api == "" {
				log.Infof("暂不支持此类型查询")
				continue
			}
			// 调用
			// data, err := httpGetAndPost(api, requestMode, b.String(), contentType, timeOut)
			data, err := httpGetAndPost(api, requestMode, b.String(), contentType, timeOut)
			if err != nil {
				log.Infof("name:%s,url:%s调用失败,err:%v", name, api, err)
				fmt.Printf("name:%s,url:%s调用失败,err:%v\n", name, api, err)
				continue
			}

			log.Infof("api===>%s,返回数据：%s\n", api, data)
			fmt.Printf("接口名称：%s\n,接口url===>%s\n,\n返回数据：%s\n\n", name, api, data)

		}
	}

}

type PersonInfo struct {
	Qlr  string `json:"qlr"`  // 权利人名字
	Sfzh string `json:"sfzh"` // 权利人身份证
}

type PersonInfo4JD struct {
	Qlrmc  string `json:"qlrmc"`  // 权利人名称
	Qlrzjh string `json:"qlrzjh"` // 权利人证件号
	Bdcqzh string `json:"bdcqzh"` // 不动产权证号
}

// 拼装请求路径以及添加参数
// 返回参数一：请求URL
// 返回参数二：请求方式
func appendUrl(name string, url string, individual Individual) map[string]string {
	m := make(map[string]string)
	switch name {
	case "房屋权属信息":
		return packageHouseOwnerUrl(individual, m)
	case "社会法人社保缴费信息":
		return packageSecurityPayment(url, individual, m)
	case "社会法人社保欠缴信息":
		return packageSecurityArrears(url, individual, m)
	case "自然人社保缴纳信息":
		return packageNaturalSecurityPayment(url, individual, m)
	case "婚姻登记信息":
		return packageMarriageInfo(url, individual, m)
	case "自然人守信红名单":
		return packageNaturalKeepPromisesList(url, individual, m)
	case "社会法人守信红名单":
		return packageSecurityKeepPromisesList(url, individual, m)
	case "自然人严重失信黑名单":
		return packageNaturalLostPromisesList(url, individual, m)
	default:
		return m
	}

	return m
}

func packageNaturalLostPromisesList(url string, individual Individual, m map[string]string) map[string]string {
	var api = UrlPrefix + "gateway/api/1/zrryzsxhmdfy"
	api = api + "?ROWNUM=0&PAGESIZE=1&SFZJHM=" + individual.Subject.Id
	log.Infof("自然人严重失信黑名单Url:", api)
	m[api] = "GET"
	return m
}

func packageSecurityKeepPromisesList(url string, individual Individual, m map[string]string) map[string]string {
	var api = UrlPrefix + "gateway/api/1/shfrcssxhmdfyfw"
	api = api + "?ROWNUM=0&PAGESIZE=1&QYMC=" + individual.Subject.Name
	log.Infof("自然人守信红名单Url:", api)
	m[api] = "GET"
	return m
}

func packageNaturalKeepPromisesList(url string, individual Individual, m map[string]string) map[string]string {
	var api = UrlPrefix + "gateway/api/1/zrrcssxhmdfy"
	api = api + "?ROWNUM=0&PAGESIZE=2&XM=" + individual.Subject.Name
	log.Infof("自然人守信红名单Url:", api)
	m[api] = "GET"
	return m
}

func packageMarriageInfo(url string, individual Individual, m map[string]string) map[string]string {
	var api = UrlPrefix + "gateway/api/1/hydjxxcx"
	api = api + "?gmsfzh=" + individual.Subject.Id
	log.Infof("婚姻登记信息Url:", api)
	m[api] = "GET"
	return m
}

func packageNaturalSecurityPayment(url string, individual Individual, m map[string]string) map[string]string {
	var api = UrlPrefix + "gateway/api/1/zrrsbjnxxcxjk"
	api = api + "?xm=" + individual.Subject.Name + "&zjhm=" + individual.Subject.Id
	log.Infof("自然人社保缴纳信息Url:", api)
	m[api] = "POST"
	return m
}

func packageSecurityArrears(url string, individual Individual, m map[string]string) map[string]string {
	var api = UrlPrefix + "gateway/api/1/shfrsbqjxxcxjk?jgqc=" + individual.Subject.Name
	log.Infof("社会法人社保欠缴信息Url:", api)
	m[api] = "GET"
	return m
}

func packageSecurityPayment(url string, individual Individual, m map[string]string) map[string]string {
	var api = UrlPrefix + "gateway/api/1/shfrsbjfxxcxjk?jgqc=" + individual.Subject.Name
	log.Infof("社会法人社保缴费信息Url:", api)
	m[api] = "GET"
	return m
}

func packageHouseOwnerUrl(individual Individual, m map[string]string) map[string]string {
	// 高邮市房屋权属查询接口
	// var p = PersonInfo{Qlr: individual.Subject.Name, Sfzh: individual.Subject.Id}
	// bytes, _ := json.Marshal(p)

	// var api = UrlPrefix + "gateway/api/1/gysfwqscxjk"
	// api = api + "?args={\"selarea\": \"\",\"clientusername\": \"1\",\"clientusercid\": \"2\",\"ytcn\": \"test\",\"computerid\": \"4\",\"computermac\": \"5\",\"computername\": \"6\",\"psw\": \"7\",\"cxrzp\": \"8\",\"qlrlist\":[" + string(bytes) + "]}"
	// log.Infof("高邮市房屋权属查询接口Url:", api)
	// m[api] = "POST"
	//
	// // 仪征市房屋权属查询接口
	// var api2 = UrlPrefix + "gateway/api/1/yzsfwqscxjk"
	// api2 = api2 + "?args={\"selarea\": \"\",\"clientusername\": \"1\",\"clientusercid\": \"2\",\"ytcn\": \"test\",\"computerid\": \"4\",\"computermac\": \"5\",\"computername\": \"6\",\"psw\": \"7\",\"cxrzp\": \"8\",\"qlrlist\":[" + string(bytes) + "]}"
	// log.Infof("仪征市房屋权属查询接口Url:", api2)
	// m[api2] = "POST"
	//
	// // 覆盖范围邗江区、广陵区、经济开发区和江都区
	// var api3 = UrlPrefix + "gateway/api/2/fwqscx"
	// api3 = api3 + "?args={\"selarea\":\"\",\"clientusername\":\"1\",\"clientusercid\":\"2\",\"ytcn\":\"test\",\"computerid\":\"4\",\"computermac\":\"5\",\"computername\":\"dsj\",\"psw\":\"7\",\"cxrzp\":\"8\",\"qlrlist\":[" + string(bytes) + "]}"
	// log.Infof("覆盖范围邗江区、广陵区、经济开发区和江都区接口Url:", api3)
	// m[api3] = "POST"

	// 覆盖范围邗江区、广陵区、经济开发区和江都区
	var api4 = UrlPrefix + "gateway/api/1/fwqscx"
	// api4 = api4 + "?args={\"selarea\":\"\",\"clientusername\":\"1\",\"clientusercid\":\"2\",\"ytcn\":\"test\",\"computerid\":\"4\",\"computermac\":\"5\",\"computername\":\"dsj\",\"psw\":\"7\",\"cxrzp\":\"8\",\"qlrlist\":[" + string(bytes) + "]}"
	log.Infof("覆盖范围邗江区、广陵区、经济开发区和江都区接口Url4:", api4)
	m[api4] = "POST"

	return m
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
