package main

import (
	"fmt"
	"time"
)

func main() {
	currentT := 1644196057
	// tm := time.Unix(int64(currentT-3600*24*8), 0)
	tm := time.Unix(int64(currentT+3600*24*22), 0)

	fmt.Println(tm)
	date := tm.AddDate(0, 1, 0)
	fmt.Println(date)
	addDate := date.AddDate(0, 0, 29)
	fmt.Println(addDate)
	t := addDate.AddDate(0, 1, 0)
	fmt.Println(t)
}
