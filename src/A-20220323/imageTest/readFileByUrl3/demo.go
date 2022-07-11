package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// imgUrl := "http://qiniu.yueda.vip/0000.jpg"
	// imgUrl := "https://horifon.oss-cn-shanghai.aliyuncs.com/20220511/f323ce8e00304625ab3cbdb61914b8e2.jpg"
	imgUrl := "http://qn.yueda.vip/yes.txt"

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
	fmt.Println("data:", string(data))

	// imageBase64 := base64.StdEncoding.EncodeToString(data)
	// fmt.Println("base64", imageBase64)

}
