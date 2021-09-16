package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 3; i++ {
			go func(ii int) {
				for j := 0; j < 3; j++ {
					fmt.Printf("i:%d,j:%d,%s\n", ii, j, "----------------")
				}
			}(i)
		}
	}()

	time.Sleep(time.Second * 2)
}
