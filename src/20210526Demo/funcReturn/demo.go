package main

import "fmt"

// 函数闭包
func adder() func(int) int { // 函数作为返回值
	sum := 0                 // 自由变量，每次调用都保留上次的结果
	return func(v int) int { // v是局部变量
		sum += v
		return sum
	}
}

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	// a := adder() is trivial and also works.
	a := adder()
	for i := 0; i < 10; i++ {
		var s int
		s = a(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, s)
	}

	f := Fibonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
