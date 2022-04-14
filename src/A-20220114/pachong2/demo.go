package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func main() {
	addressList := []string{
		"0x00000000219ab540356cbb839cbe05303d7705fa",
		// "0x00000000219ab540356cbb839cbe05303d7705fa",
		// "0x00000000219ab540356cbb839cbe05303d7705fb",
		// "0x00000000219ab540356cbb839cbe05303d7705fc",
	}
	for _, addr := range addressList {
		fmt.Println("addr:", addr)
		// requestUrl := "https://www.liaoxuefeng.com/"
		// https://dune.xyz/labels/ethereum/0x00000000219ab540356cbb839cbe05303d7705fa
		requestUrl := "https://dune.xyz/labels/ethereum/"
		requestUrl = requestUrl + addr
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
		fmt.Println("content:", content)

		// 下面主要是解析标签
		doc := soup.HTMLParse(content)
		subDocs := doc.FindAll("table", "class", "table_table__fuS_N")
		fmt.Println("===>", subDocs)
		fmt.Println("-------------------------------------------------------------")
		for _, subDoc := range subDocs {
			link := subDoc.Find("td")
			fmt.Println(link.Text())
		}
	}

}
