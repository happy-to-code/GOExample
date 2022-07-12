package main

import (
	"fmt"
	"time"
)

func main() {
	var sList = []string{"foo", "bar", "baz", "abc"}
	for _, s := range sList {
		fmt.Printf("value: %s\n", s)
	}

	now := time.Now()
	// fmt.Println("s==>", now.Unix())
	fmt.Println("ms==>", now.UnixMilli())
	// fmt.Println("ws==>", now.UnixMicro())
	fmt.Println("================================================================")
	msStr := fmt.Sprintf("%d", now.UnixMilli())
	s := msStr[10:]
	fmt.Println("333:", s)

	// 1656557388276
	fmt.Println("===========now:", now)
	timeStr := now.Format("20060102150405")
	fmt.Println("timeStr==>", timeStr)

}
