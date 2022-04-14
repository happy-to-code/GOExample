package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func main() {
	requestUrl := "https://www.liaoxuefeng.com/"
	// 发送Get请求
	rsp, err := http.Get(requestUrl)
	if err != nil {
		log.Println(err.Error())
		return
	}
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	content := string(body)
	defer rsp.Body.Close()

	// 下面主要是解析标签
	doc := soup.HTMLParse(content)
	subDocs := doc.FindAll("div", "class", "uk-margin")
	for _, subDoc := range subDocs {
		link := subDoc.Find("a")
		fmt.Println(link.Text())
	}
}
