package main

import (
	"fmt"
	"time"
)

func main() {
	var eventBytes string

	defer func() {
		fmt.Println("defer:", eventBytes)
	}()

	time.Sleep(time.Second * 1)
	eventBytes = "hello"
	fmt.Println("end")
}
