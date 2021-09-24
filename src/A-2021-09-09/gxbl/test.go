package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var lock sync.Mutex

	var errNum uint8

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			lock.Lock()
			errNum++
			lock.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println("-->", errNum)
}
