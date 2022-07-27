package main

import (
	"fmt"
	"time"
)

func main() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
	tick1 := time.NewTicker(1 * time.Second)
	// 初始化channel个数
	ch := make(chan struct{}, 300)
	for {
		select {
		case <-tick1.C:
			for i := 0; i < 100000; i++ {
				// 这里往channel写数据
				ch <- struct{}{}
				go func() {
					if err := recover(); err != nil {
						fmt.Println(err)
					}
					fmt.Println("xxx==>", time.Now())
					time.Sleep(10 * time.Second)
					// 读取channel数据
					<-ch
				}()
			}
		}
	}
	time.Sleep(2 * time.Hour)
}
