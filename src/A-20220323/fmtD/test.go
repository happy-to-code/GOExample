package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// fmt.Println(1)
	// fmt.Println(time.Now().Unix())
	// fmt.Println(time.Now())
	//
	// now := time.Now()
	// nowMulOneDay := now.AddDate(0, 0, -1)
	// str := timeStamp2Str(nowMulOneDay.Unix())
	// fmt.Println(str)
	// fmt.Println(strings.HasPrefix("0x123", "0x"))

	fmt.Println("================================")
	jwtTimeOut := 48
	now := time.Now()
	expireTime := now.Add(time.Duration(jwtTimeOut) * time.Hour)
	fmt.Println(expireTime)
	fmt.Println(expireTime.Unix())

	fmt.Println(expireTime.Unix() - now.Unix())

	var s string
	s = " "
	fmt.Println(len(strings.TrimSpace(s)))

}

const timeTemplate = "2006-01-02 15:04:05"

func timeStamp2Str(t int64) string {
	return time.Unix(t, 0).Format(timeTemplate)
}
