package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var lk sync.Mutex

	var sList []string

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()
			for j := 1; j <= 2; j++ {
				lk.Lock()
				sList = append(sList, fmt.Sprintf("go-fun-%d-%d", index, j))
				lk.Unlock()
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("len():%d,\nsList:%v\n", len(sList), sList)

	var ch = make(chan int, 100)
	for i := 0; i < 100; i++ {
		ch <- i
	}

	for k := range ch {
		fmt.Println("k:", k)
		if k == 99 {
			close(ch)
		}
	}
}
