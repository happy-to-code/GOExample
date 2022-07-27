package main

import (
	"fmt"
	"time"
)

var c = make(chan int, 1)

func main() {

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("i: ", i)
			c <- 100
		}
	}()
	go func() {
		for j := 0; j < 5; j++ {
			fmt.Println("j: ", j)
			c <- 200
		}
	}()

	go func() {
		time.Sleep(time.Second * 2)
		for {
			select {
			case num := <-c:
				fmt.Println("num:", num)
			}
		}
	}()

	time.Sleep(time.Second * 3)
}
