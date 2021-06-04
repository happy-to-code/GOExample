package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"strconv"
)

type PostDemo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// http://localhost:9088/v1/api/world
	// url := "http://localhost:9088/v1/api/world"
	// for i := 0; i < 10; i++ {
	// 	err, s := httpGet(url)
	// 	if err == nil {
	// 		// panic(err)
	// 		fmt.Println(s)
	// 	}
	// }

	// 	========================================================
	url := "http://localhost:9088/v1/api/post"
	for i := 0; i < 10; i++ {
		var p PostDemo
		p.Age = i
		p.Name = "aa" + strconv.Itoa(i)
		jsonBytes, _ := json.Marshal(p)
		// fmt.Println("---->", string(jsonBytes))
		err, s := httpPost(url, string(jsonBytes))
		if err == nil {
			fmt.Println(s)
		}
	}
}

func httpPost(url, jsonData string) (error, string) {
	req := &fasthttp.Request{}
	req.SetRequestURI(url)

	requestBody := []byte(jsonData)
	req.SetBody(requestBody)

	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	resp := &fasthttp.Response{}

	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		fmt.Println("请求失败：", err)
		return err, ""
	}

	b := resp.Body()
	// fmt.Println("result:\r\n", string(b))
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

func main1() {
	url := `http://www.baidu.com`

	status, resp, err := fasthttp.Get(nil, url)
	if err != nil {
		fmt.Println("请求失败:", err.Error())
		return
	}

	if status != fasthttp.StatusOK {
		fmt.Println("请求没有成功:", status)
		return
	}

	fmt.Println(string(resp))
}
func main2() {
	url := `http://httpbin.org/post?key=123`

	// 填充表单，类似于net/url
	args := &fasthttp.Args{}
	args.Add("name", "test")
	args.Add("age", "18")

	status, resp, err := fasthttp.Post(nil, url, args)
	if err != nil {
		fmt.Println("请求失败:", err.Error())
		return
	}

	if status != fasthttp.StatusOK {
		fmt.Println("请求没有成功:", status)
		return
	}

	fmt.Println(string(resp))
}
func main3() {
	url := `http://httpbin.org/post?key=123`

	req := &fasthttp.Request{}
	req.SetRequestURI(url)

	requestBody := []byte(`{"request":"test"}`)
	req.SetBody(requestBody)

	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	resp := &fasthttp.Response{}

	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		fmt.Println("请求失败:", err.Error())
		return
	}

	b := resp.Body()

	fmt.Println("result:\r\n", string(b))
}
func main4() {
	url := `http://httpbin.org/post?key=123`

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		// 用完需要释放资源
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	req.SetRequestURI(url)

	requestBody := []byte(`{"request":"test"}`)
	req.SetBody(requestBody)

	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Println("请求失败:", err.Error())
		return
	}

	b := resp.Body()

	fmt.Println("result:\r\n", string(b))
}
