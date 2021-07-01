package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	fmt.Println("t:", t)

	var a int64 = 10
	fmt.Println(time.Second * 10)
	fmt.Println(time.Duration(10 * 1000000000))
	fmt.Println(time.Duration(a) * time.Second)
	fmt.Println("-------------------------------")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
