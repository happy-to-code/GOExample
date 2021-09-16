package main

import "time"

func main() {
	var t int64 = 1631519025232
	var s = time.Unix(t/1000, 0).Format("2006-01-02 15:04:05")
	println(s)
}
