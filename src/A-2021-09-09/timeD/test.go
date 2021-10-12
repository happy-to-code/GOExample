package main

import (
	"fmt"
	"time"
)

func main() {
	var s = "20210908"

	loc, _ := time.LoadLocation("Local")
	the_time, err := time.ParseInLocation("20060102", s, loc)
	if err != nil {
		panic(err)
	}
	fmt.Println(the_time)
	fmt.Println(the_time.Unix())
	timeStr := the_time.Format("2006-01-02")
	fmt.Println(timeStr)
}
