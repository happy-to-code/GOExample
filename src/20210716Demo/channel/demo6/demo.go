package main

import (
	"fmt"
	"time"
)

var syncch = make(chan int, 100)

func sendSyncMsg(smsg int) {
	syncch <- smsg
}

func main() {

	go dealMsg(syncch)

	for i := 1; i <= 20; i++ {
		time.Sleep(time.Millisecond * 100)
		sendSyncMsg(i)
	}

	time.Sleep(2)
}

func dealMsg(c chan int) {
	for {
		for i := range c {
			fmt.Println("dealMsg data:", i)
		}
	}
}
