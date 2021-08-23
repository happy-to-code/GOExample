package main

import "A-2021-08-11/cycle/p1"

// 测试：包的循环引用
func main() {
	var h p1.SayHello
	h.SayHello()       // i'm p1, hello!!!
	h.SayHelloFromP2() // i'm p1, hello!!!
}
