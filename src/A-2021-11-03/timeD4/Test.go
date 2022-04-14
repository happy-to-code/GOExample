package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	var t = "1640586228435"
	stamp := timeStr2TimeStamp(t)
	fmt.Println(stamp)

	// 1640587033
	// 1640586124

}

const timeLayout = "2006-01-02 15:04:05"

// 日期转换成时间戳
func timeStr2TimeStamp(timeStr string) int64 {
	now := time.Now().Unix()
	if timeStr = strings.TrimSpace(timeStr); timeStr == "" {
		return now
	}
	timeInt64, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		return now
	}
	return timeInt64 / 1000

	// location, err := time.LoadLocation("Local")
	// if err != nil {
	// 	return now
	// }
	// parse, err := time.ParseInLocation(timeLayout, timeStr, location)
	// if err != nil {
	// 	return now
	// }
	// return parse.Unix()
}
