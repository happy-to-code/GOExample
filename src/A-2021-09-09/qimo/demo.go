package main

import (
	"fmt"
	"time"
)

func main() {
	var list []int

	var i = 100
	mod := i % 100
	if mod == 0 && i > 100 {
		for j := i - 199; j <= i-100; j++ {
			list = append(list, j)
		}
	}

	fmt.Println("list:", list)
	fmt.Println("list len:", len(list))
	fmt.Println(int(time.Now().Unix()))
}
