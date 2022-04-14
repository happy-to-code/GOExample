package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	var unix = time.Now().AddDate(0, 0, 1).Unix()
	timeEnd := time.Unix(unix, 0).Format("2006-01-02 15:04:05")
	fmt.Println(timeEnd)

	now := time.Now()

	yesterday := now.AddDate(0, 0, -1)

	last7Days := now.AddDate(0, 0, -7)

	fmt.Println(now)
	fmt.Println(yesterday)
	fmt.Println(last7Days)
	fmt.Println(strconv.FormatInt(int64(18000000000000), 16))

	var ss = "0x123456789"
	fmt.Println(ss[2:10])

	var i int64 = 100000000000000
	s := strconv.FormatInt(i, 16)
	fmt.Println(s)

	fmt.Println(strings.HasPrefix("0x2557ee3", "0x"))

}
