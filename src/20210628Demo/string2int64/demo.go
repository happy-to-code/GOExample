package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	s := "1625035954079"
	atoi, _ := strconv.Atoi(s)
	fmt.Println(atoi / 1000)
	fmt.Println(atoi/1000 + 1)
	fmt.Println(int(time.Now().Unix()))
}
