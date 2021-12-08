package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 50)
	go chPrint(ch)

	for i := 0; i < 100; i++ {
		send2Chan(ch, i)
	}
	fmt.Println("end^")
	time.Sleep(time.Second * 3)
}

func send2Chan(ch chan int, i int) {
	ch <- i
	// select {
	// case ch <- i:
	// default:
	// 	log.Infof("%s", "failed to submit message ")
	// 	fmt.Println("-----------",i)
	// }
	// return
}

func chPrint(ch chan int) {
	time.Sleep(time.Second * 3)
	for i := range ch {
		// if i%10 == 0 {
		// 	time.Sleep(time.Second * 1)
		// }
		fmt.Println(i)
	}
}
