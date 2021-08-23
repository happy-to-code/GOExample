package main

import "fmt"

func main() {

	m := make(map[string]uint64, 3)
	// m["a"] = 2
	// m["b"] = 3
	// m["c"] = 0

	fmt.Println(m["a"])
	fmt.Println(m["aa"])
}
