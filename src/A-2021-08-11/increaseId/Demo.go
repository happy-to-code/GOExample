package main

import (
	"fmt"
	"reflect"
	"time"
)

type AutoInc struct {
	start, step int64
	queue       chan int64
	running     bool
}

func New(start, step int64) (ai *AutoInc) {
	ai = &AutoInc{
		start:   start,
		step:    step,
		running: true,
		queue:   make(chan int64, 4),
	}
	go ai.process()
	return
}

func (ai *AutoInc) process() {
	defer func() { recover() }()
	for i := ai.start; ai.running; i = i + ai.step {
		ai.queue <- i
	}
}

func (ai *AutoInc) Id() int64 {
	return <-ai.queue
}

func (ai *AutoInc) Close() {
	ai.running = false
	close(ai.queue)
}

func main() {
	var ai *AutoInc
	ai = New(1, 1)

	for {
		a := ai.Id()
		time.Sleep(1 * time.Second)

		fmt.Println(a, reflect.TypeOf(a))
		if a == 5 {
			break
		}
	}
	fmt.Println("=============")
	ai = New(0, 1)

	for {
		a := ai.Id()
		time.Sleep(1 * time.Second)

		fmt.Println(a, reflect.TypeOf(a))
		if a == 5 {
			break
		}
	}

}
