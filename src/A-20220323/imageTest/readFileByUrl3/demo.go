package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	imgUrl := "http://qiniu.yueda.vip/0000.jpg"

	// 获取远端图片
	res, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	defer res.Body.Close()

	// 读取获取的[]byte数据
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	imageBase64 := base64.StdEncoding.EncodeToString(data)
	fmt.Println("base64", imageBase64)

}
