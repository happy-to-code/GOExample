package main

import "fmt"

var eventMsgHeightChan = make(chan uint64, 5000)

func main() {
	go dealEventMsgHeightByBlockHeight()
	for i := 10; i < 90; i++ {
		eventMsgHeightChan <- uint64(i)
	}
	for {

	}
}

func dealEventMsgHeightByBlockHeight() {
	for {
		select {
		case height := <-eventMsgHeightChan:
			fmt.Println("dealEventMsgHeightByBlockHeight---------------->", height)
			var err error
			if height%2 == 0 {
				err = fmt.Errorf("我是错误")
			}
			if err != nil {
				fmt.Println("err:", err)
				continue
			}
			fmt.Println("last:", height)
		}
	}
}
