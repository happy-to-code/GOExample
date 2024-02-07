package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	chanA := make(chan struct{}, 1)
	chanB := make(chan struct{}, 1)
	chanC := make(chan struct{}, 1)

	wg.Add(9)
	chanA <- struct{}{}

	for i := 0; i < 3; i++ {

		go printB(chanB, chanC, &wg)
		go printA(chanA, chanB, &wg)
		go printC(chanC, chanA, &wg)
	}

	wg.Wait()
	fmt.Println("finish")
}

func printA(start, end chan struct{}, wg *sync.WaitGroup) {
	<-start
	fmt.Println("A")
	wg.Done()
	end <- struct{}{}
}
func printB(start, end chan struct{}, wg *sync.WaitGroup) {
	<-start
	fmt.Println("B")
	wg.Done()
	end <- struct{}{}
}
func printC(start, end chan struct{}, wg *sync.WaitGroup) {
	<-start
	fmt.Println("C")
	wg.Done()
	end <- struct{}{}
}
