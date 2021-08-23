package p2

import (
	"fmt"
)

type ISayHello interface {
	SayHello()
}

type p2UseP1 struct {
	realP1 ISayHello
}

// New 如果不能直接使用p1，就把p1传过来，我给个interface接收下，我在内部使用
func New(h ISayHello) *p2UseP1 {
	return &p2UseP1{
		h,
	}
}

// SayP1Hello p1 say hello
func (p *p2UseP1) SayP1Hello() {
	p.realP1.SayHello()
}

// SayP2Hello p2 say hello
func (p *p2UseP1) SayP2Hello() {
	fmt.Println("i'm p2, hello!!!")
}
