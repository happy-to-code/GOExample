package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()

	type List []int
	var list []List

	maxBlockNumber := 10

	var l []int
	for i := 1; i <= maxBlockNumber; i++ {
		l = append(l, i)
		if i%20 == 0 {
			// fmt.Println("-->", i)
			list = append(list, l)

			l = nil
		} else {
			if l != nil && i == maxBlockNumber {
				list = append(list, l)
			}
		}
	}
	time.Sleep(time.Second)
	fmt.Println(list)
	t2 := time.Now()

	fmt.Println(t2.Sub(t1).String())
	fmt.Println(t2.Sub(t1).String())
}
