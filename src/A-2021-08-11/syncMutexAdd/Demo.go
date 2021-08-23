package main

import "sync"

var t TotalAmount
var wg sync.WaitGroup

func main() {
	t := &TotalAmount{}
	for i := 0; i < 9999; i++ {
		wg.Add(1)
		go mutexPlus(t)
	}

	wg.Wait()
	println(t.Amount)
}

type TotalAmount struct {
	Amount int64
	sync.Mutex
}

func mutexPlus(t *TotalAmount) {
	defer wg.Done()
	t.Lock()
	t.Amount++
	t.Unlock()
}
