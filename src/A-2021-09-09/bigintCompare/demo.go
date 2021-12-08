package main

import (
	"fmt"
	"math"
)

func main() {
	// 18446744073709551615
	// var max uint64 = 18446744073719
	var max uint64 = 140000000000000
	var maxUint64 uint64 = math.MaxUint64
	fmt.Println(maxUint64)
	fmt.Println(maxUint64 / 1000000)

	println(compare(max, 6))
}

// a要比较的值,b精度
func compare(a, b uint64) bool {
	return a > math.MaxUint64/1000000
}
