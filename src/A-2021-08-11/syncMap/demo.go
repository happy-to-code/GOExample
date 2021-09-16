package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	m.Store("k1", 16)
	m.Store("k2", 6)
	m.Store("k3", 66)

	load, ok := m.Load("k1")
	fmt.Println(load, ok)
	fmt.Printf("%T\n", load)
	fmt.Println("=========================")
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("%T\n", key)
		fmt.Println(fmt.Sprintf("%s", key))
		fmt.Println("Key =", key, "Value =", value)
		return true
	})
}

func walk(key, value interface{}) bool {
	fmt.Println("Key =", key, "Value =", value)
	return true
}
