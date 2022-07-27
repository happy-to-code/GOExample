package main

import "fmt"

func main() {
	a := 4
	b := 2
	fmt.Println("add:", f(a, b, add))
	fmt.Println("sub:", f(a, b, sub))
	fmt.Println("mul:", f(a, b, mul))
	fmt.Println("div:", f(a, b, div))

}

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}
func div(a, b int) int {
	return a / b
}

func f(a, b int, ff func(int, int) int) int {
	return ff(a, b)
}
