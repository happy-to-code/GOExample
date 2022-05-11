package main

import (
	"fmt"
	"strings"
)

func main() {
	imgUrl := "http://qiniu.yueda.vip/0000.jpg"
	imgUrl2 := "http://qiniu.yueda.vip/0000.png"
	fmt.Println(imgUrl)

	println(strings.HasSuffix(imgUrl, ".jpg"))
	println(strings.HasSuffix(imgUrl2, ".png"))

}
