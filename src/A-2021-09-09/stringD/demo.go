package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "*|扬州金泉旅游用品股份有限公司|扬州金泉旅游用品股份有限公司|扬州金泉旅游用品股份有限公司|扬州金泉旅游用品股份有限公司|扬州金泉旅游用品股份有限公司|扬州金泉旅游用品股份有限公司|扬州金泉旅游用品股份有限公司|扬州金泉旅游用品股份有限公司|扬州金泉旅游用品股份有限公司|扬州金泉旅游用品股份有限公司|扬州金泉旅游用品股份有限公司"
	// var s = ""
	split := strings.Split(strings.Replace(strings.TrimSpace(s), "*", "", 1), "|")
	fmt.Println(split[0])
	fmt.Println(split)
	fmt.Println(len(split))
	fmt.Println(split[len(split)-1])

	var countFromAddr string
	fmt.Println(countFromAddr == "")

}
