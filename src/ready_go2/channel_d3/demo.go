package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println(runtime.GOMAXPROCS)

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(a int) {
			defer wg.Done()
			fmt.Println("=====>", a)
		}(i)
	}
	wg.Wait()
}
