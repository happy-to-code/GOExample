package p1

import (
	"A-2021-08-11/cycle/p2"
	"fmt"
)

type SayHello struct{}

func (*SayHello) SayHello() {
	fmt.Println("i'm p1, hello!!!")
}

func (h *SayHello) SayHelloFromP2() {
	pup := p2.New(h)
	pup.SayP1Hello()
}
