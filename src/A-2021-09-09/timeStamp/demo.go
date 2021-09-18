package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	var t int64 = 1631519025232
	var s = time.Unix(t/1000, 0).Format("2006-01-02 15:04:05")
	println(s)

	fmt.Println("--------------")
	var tStr = "2021-09-16 14:42:32"

	fmt.Println("==>", time.Now().Unix())
	fmt.Println("=--=>", timeStr2TimeStamp(tStr))

	host, port, err := net.SplitHostPort("http://10.1.3.150:9090")
	if err != nil {
		panic(err)
	}

	fmt.Println(host, "--", port)

}

// 日期转换成时间戳
func timeStr2TimeStamp(timeStr string) int64 {
	now := time.Now().Unix()
	if timeStr = strings.TrimSpace(timeStr); timeStr == "" {
		return now
	}
	location, err := time.LoadLocation("Local")
	if err != nil {
		return now
	}
	parse, err := time.ParseInLocation(timeLayout, timeStr, location)
	if err != nil {
		return now
	}
	return parse.Unix()
}

const (
	timeLayout = "2006-01-02 15:04:05"
)
