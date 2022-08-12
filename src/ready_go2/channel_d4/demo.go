package main

import "fmt"

var ch chan int

func elegance() {
	<-ch
	fmt.Println("the ch value receive", ch)
}

func main() {
	ch = make(chan int, 5) // 控制goroutine数量
	for i := 0; i < 100; i++ {
		ch <- 1
		fmt.Println("the ch value send", ch)
		go elegance()
		fmt.Println("the result i", i)
	}

}
