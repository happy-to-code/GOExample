package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// var s = "-1"
	//
	// fmt.Println(convertStr2Num(s)+1)
	//
	// ss := "0xdAC17F958D2ee523a2206206994597C13D831ec7"
	// fmt.Println(strings.ToLower(ss))

}
func convertStr2Num(s string) int64 {
	if s = strings.TrimSpace(s); s == "" {
		return 0
	}
	parseInt, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Printf("[%s]转换失败:%v", s, err)
		return 0
	}
	return parseInt
}
