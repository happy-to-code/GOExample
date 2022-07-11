package main

//
// import (
// 	"fmt"
// 	"github.com/valyala/fasthttp"
// 	// "github.com/buaazp/fasthttprouter"
// 	"time"
// )
//
// func main() {
// 	var arg = &fasthttp.Args{}
// 	arg.Set("hello", "world")
// 	fmt.Println(doTimeout(arg, "GET", "http://127.0.0.1/hello", nil))
// 	fmt.Println(doJsonTimeout("GET", "http://127.0.0.1/hello", "{\"hello\":\"world\"}"))
//
// }
//
// // http请求
// func doTimeout(arg *fasthttp.Args, method string, requestURI string, cookies map[string]interface{}) ([]byte, int, error) {
// 	req := &fasthttp.Request{}
// 	switch method {
// 	case "GET":
// 		req.Header.SetMethod(method)
// 		// 拼接url
// 		requestURI = requestURI + "?" + arg.String()
// 	case "POST":
// 		req.Header.SetMethod(method)
// 		arg.WriteTo(req.BodyWriter())
// 	}
// 	if cookies != nil {
// 		for key, v := range cookies {
// 			req.Header.SetCookie(key, v.(string))
// 		}
// 	}
// 	req.SetRequestURI(requestURI)
//
// 	resp := &fasthttp.Response{}
// 	err := gCli.DoTimeout(req, resp, time.Second*30)
//
// 	return resp.Body(), resp.StatusCode(), err
// }
// func doJsonTimeout(method string, url, bodyjson string) ([]byte, int, error) {
// 	req := &fasthttp.Request{}
// 	resp := &fasthttp.Response{}
//
// 	switch method {
// 	case "GET":
// 		req.Header.SetMethod(method)
// 	case "POST":
// 		req.Header.SetMethod(method)
// 	}
//
// 	req.Header.SetContentType("application/1.json")
// 	req.SetBodyString(bodyjson)
//
// 	req.SetRequestURI(url)
//
// 	err := gCli.DoTimeout(req, resp, time.Second*30)
// 	return resp.Body(), resp.StatusCode(), err
// }
