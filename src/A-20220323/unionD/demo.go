package main

import (
	"fmt"
	"time"
)

func main() {
	var s1 = []string{"1", "2", "3", "5"}
	var s2 = []string{"4"}
	fmt.Println(union(s1, s2))

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}

// 求并集
func union(slice1, slice2 []string) []string {
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}
