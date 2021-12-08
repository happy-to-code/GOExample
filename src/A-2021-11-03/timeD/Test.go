package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	fmt.Println(now.UnixNano() / 1e6)

	date := now.AddDate(0, 3, 0)
	fmt.Println(date)

	// 2021-11-10-17-04-18-90604afc0aaaccd848597d9d362e871da6906af3b810cbf9ba00d8e326a38db2
	// 当前时间字符串
	timeStr := time.Unix(time.Now().Unix(), 0).Format("2006-01-02-15-04-05")
	hash := "90604afc0aaaccd848597d9d362e871da6906af3b810cbf9ba00d8e326a38db2"
	itemName := "婚姻数据"

	hashHead := hash[:6]
	hashEnd := hash[len(hash)-6:]
	fmt.Println(hashHead, " ", hashEnd)

	fileName := timeStr + "-" + itemName + "-" + fmt.Sprintf("%s%s%s", hash[:6], "...", hash[len(hash)-6:])
	fmt.Println(fileName)
}
