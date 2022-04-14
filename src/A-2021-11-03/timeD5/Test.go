package main

import (
	"fmt"
)

func main() {
	str1 := "微信公众号菜鸟\n干货多多"
	fmt.Println(str1)
	fmt.Println("----------------------------------------")

	str1 = "微信公众号菜鸟\r干货多多"
	fmt.Println(str1)
	fmt.Println("----------------------------------------")

	str1 = "微信公众号菜鸟\r\n干货多多"
	fmt.Println(str1)
	fmt.Println("----------------------------------------")

}
