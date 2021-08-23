package main

import "fmt"

func main() {
	ch := make(chan int) // 创建int 型的无缓冲的channel
	go func() {
		sum := 0
		for i := 0; i < 100; i++ {
			sum += i
		}
		ch <- sum // 将goroutine 算出的数放进通道
	}()
	fmt.Println(<-ch) // 从通道获取数据，退出main函数，如果channel还没有存入数据，就会阻塞等待
}
