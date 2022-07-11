package main

import (
	"fmt"
	"time"
)

func main() {
	// imgUrl := "http://qiniu.yueda.vip/0000.jpg"
	// imgUrl2 := "http://qiniu.yueda.vip/0000.png"
	// fmt.Println(imgUrl)
	//
	// println(strings.HasSuffix(imgUrl, ".jpg"))
	// println(strings.HasSuffix(imgUrl2, ".png"))

	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("len:", len(list))
	fmt.Println("cap:", cap(list))
	fmt.Println("==>list1", list)

	fmt.Println("-----------------------------")

	list2 := list[3:6]
	fmt.Println("len:", len(list2))
	fmt.Println("cap:", cap(list2))
	fmt.Println("==>list2", list2)

	fmt.Println(time.Now().Unix())
	// 1653027463
}
