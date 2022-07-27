package main

import (
	"fmt"
	"sync"
	"time"
)

var count int32

func main() {
	t := time.Now()
	fmt.Println(t)
	now := t.Add(-time.Duration(60) * time.Second)
	fmt.Println(now)

	var wg sync.WaitGroup
	var lock sync.Mutex

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			count++

			// atomic.AddInt32(&count, 1)
		}()
	}
	wg.Wait()
	fmt.Println("count:", count)

}
