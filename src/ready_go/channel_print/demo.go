package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(12) // 打印四组，三个goroutine需要执行4次，3*4表示处于等待goroutine的数量

	dogChan := make(chan struct{}, 1) // 此处如果申明的是无缓冲的通道，那么会在16行代码处处于阻塞状态！原因：通道中无数据，向通道写数据，但无协程读取。
	catChan := make(chan struct{}, 1)
	fishChan := make(chan struct{}, 1)

	dogChan <- struct{}{}
	for i := 0; i < 4; i++ { // 打印四组，三个goroutine需要执行4次
		go dog(&wg, dogChan, catChan)
		go cat(&wg, catChan, fishChan)
		go fish(&wg, fishChan, dogChan)
	}

	wg.Wait()
	fmt.Println("main finished!")
}

func dog(wg *sync.WaitGroup, dogChan chan struct{}, catChan chan struct{}) {
	<-dogChan
	fmt.Println("dog")
	wg.Done()
	catChan <- struct{}{}
}

func cat(wg *sync.WaitGroup, catChan chan struct{}, fishChan chan struct{}) {
	<-catChan
	fmt.Println("cat")
	wg.Done()
	fishChan <- struct{}{}
}

func fish(wg *sync.WaitGroup, fishChan chan struct{}, dogChan chan struct{}) {
	<-fishChan
	fmt.Println("fish")
	wg.Done()
	dogChan <- struct{}{}

}
