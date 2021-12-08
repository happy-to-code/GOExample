package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Unix())
	addOneMonth := now.AddDate(0, 1, 0)
	fmt.Println(addOneMonth.Unix())
}
