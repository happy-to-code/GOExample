package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 20)
	go func(cc chan int) {
		for k := range cc {
			fmt.Println(k, "---------")
		}
	}(c)

	for i := 0; i < 10; i++ {
		c <- i
	}

	time.Sleep(time.Second * 2)
}
