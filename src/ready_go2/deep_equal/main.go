package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a []string
	b := []string{}
	fmt.Println("a:", a, a == nil)
	fmt.Println("b:", b, b == nil)

	fmt.Println(reflect.DeepEqual(a, []string{}))
	fmt.Println(reflect.DeepEqual(b, []string{}))
	// Output:
	// false
	// true
}
