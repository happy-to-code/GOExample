package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	t := time.Now().Unix()
	fmt.Println("--->", t)
	itoa := strconv.Itoa(int(t))
	fmt.Printf("---->%T,,,%v\n", itoa, itoa)

	var s string

	fmt.Println(s == "")
}
