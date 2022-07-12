package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	dogChan := make(chan struct{}, 1)
	catChan := make(chan struct{}, 1)
	fishChan := make(chan struct{}, 1)

	dogChan <- struct{}{}
	wg.Add(15)
	for i := 0; i < 5; i++ {
		go dog(dogChan, catChan, &wg)
		go cat(catChan, fishChan, &wg)
		go fish(fishChan, dogChan, &wg)
	}
	wg.Wait()
}

func dog(receiver chan struct{}, response chan struct{}, wg *sync.WaitGroup) {
	<-receiver
	fmt.Println("1、dog")
	wg.Done()
	response <- struct{}{}
}

func cat(receiver chan struct{}, response chan struct{}, wg *sync.WaitGroup) {
	<-receiver
	fmt.Println("2、cat")
	wg.Done()
	response <- struct{}{}
}

func fish(receiver chan struct{}, response chan struct{}, wg *sync.WaitGroup) {
	<-receiver
	fmt.Println("3、fish")
	wg.Done()
	response <- struct{}{}
}
