package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	var n1 []interface{}
	n2 := []int{1, 2, 3}
	n3 := append(n1, n2)
	fmt.Println(n3)
	fmt.Println(len(n3))

	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
	fmt.Printf("%v\n", t2.Sub(t1))
}
