package main

import (
	"fmt"
	"time"
)

func main() {
	var timeList = []string{"2021-11-01 21:19:25",
		"2021-11-01 21:25:16",
		"2021-11-01 22:19:44",
		"2021-11-02 05:07:57",
		"2021-11-02 06:57:38",
		"2021-11-02 07:02:35",
		"2021-11-02 07:02:35",
		"2021-11-02 07:06:06",
		"2021-11-02 01:11:21",
		"2021-11-02 01:11:21",
		"2021-11-02 02:57:36",
		"2021-11-02 03:25:34"}
	for _, s := range timeList {
		// fmt.Println(s)
		loc, _ := time.LoadLocation("Local")
		the_time, err := time.ParseInLocation("2006-01-02 15:04:05", s, loc)
		if err == nil {
			unix_time := the_time.Unix() // 1504082441
			fmt.Println(unix_time)
		}
	}

}
