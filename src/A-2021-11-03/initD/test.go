package main

import "fmt"

func main() {
	fmt.Println("main")
}

var a = 10

func init() {
	fmt.Println(a + 1)
}
func init() {
	fmt.Println(a)
}

const b = 11

func init() {
	fmt.Println(a + 1)
}
