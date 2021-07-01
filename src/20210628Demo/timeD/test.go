package main

import (
	"fmt"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

func main() {
	t1 := "2021-07-01 12:04:05"
	loc, _ := time.LoadLocation("Local")

	the_time, err := time.ParseInLocation(TimeFormat, t1, loc)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("t1:::", the_time.Unix())
	}

	now := time.Now().Unix()
	fmt.Println("now:", now)

}
