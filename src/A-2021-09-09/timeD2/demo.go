package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// 参数的第一个值
	month, _ := time.Parse("200601", "202108")
	addOneMonth := month.AddDate(0, 1, 0)

	fmt.Println("now:", now.Year(), " ", now.Month())
	fmt.Println("addOneMonth:", addOneMonth.Year(), " ", addOneMonth.Month())

	fmt.Println(now.Month() > addOneMonth.Month())

	fmt.Println("--------------")
	var s []string = []string{"aa", "bb", "cc", "dd"}

	var s2 []string
	if len(s) > 2 {
		s2 = s[:2]
	}

	fmt.Println(s)
	fmt.Println(s2)
}
