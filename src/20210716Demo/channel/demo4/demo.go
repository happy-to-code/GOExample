package main

import (
	"fmt"
	"time"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func f2(in chan int) {
	in <- 2
}

func main() {
	out := make(chan int)
	go f2(out)
	go f1(out)
	time.Sleep(1e9)
}
