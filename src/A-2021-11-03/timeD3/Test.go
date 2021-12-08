package main

import (
	"fmt"
	"time"
)

func main() {
	var times = "2021-01-28 21:19:25"

	loc, _ := time.LoadLocation("Local")
	the_time, err := time.ParseInLocation("2006-01-02 15:04:05", times, loc)
	if err != nil {
		panic(err)
	}
	fmt.Println(the_time.Unix())

	date := the_time.AddDate(0, 1, 0)
	fmt.Println(date.Unix())
	fmt.Println("-------------")
	fmt.Println("==>", time.Now().Unix())

}
