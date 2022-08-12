package main

import (
	"fmt"
	"sync"
)

func main() {
	var msg = Msg{Name: "yida"}
	msg.Lock()
	defer msg.Unlock()
	fmt.Println(msg.Name)
}

type Msg struct {
	sync.Mutex
	Name string
}
