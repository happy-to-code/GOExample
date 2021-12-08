package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(test(uint(i)))
	// }
	result := -1 << 31
	fmt.Println(result)
}

func test(n uint) uint {
	if n == 0 || n == 1 {
		return n
	}
	return test(n-1) + test(n-2)
}
