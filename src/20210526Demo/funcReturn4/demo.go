package main

type CalculateFunc func(a, b int) int

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Calculate(a, b int, calculateFunc CalculateFunc) int {
	return calculateFunc(a, b)
}

func simple03() {
	a, b := 5, 3
	// 声明一个函数
	f := Add
	// 相加
	add := Calculate(a, b, f)
	// 相减
	sub := Calculate(a, b, Sub)

	println(add, sub)
}

func main() {
	simple03()
}
