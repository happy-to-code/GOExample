package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int32

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// count++
			atomic.AddInt32(&count, 1)
		}()
	}
	wg.Wait()

	fmt.Println("count:", count)
}
