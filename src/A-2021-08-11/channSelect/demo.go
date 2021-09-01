package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 10)
	defer close(ch)

	for i := 0; i < 5; i++ {
		fmt.Println("cap:", cap(ch))
		ch <- i
		fmt.Println("---->", len(ch))
	}

	for i := range ch {
		fmt.Println("=========================>", i)
	}

}

func sendToChan(c chan int) {
	fmt.Println("=======================")
	select {
	case c <- 1:
	case t := <-c:
		fmt.Println("value:", t)
	default:
		fmt.Println("failed")
	}
}
