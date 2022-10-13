package main

import "fmt"

// 求交集
func intersect(slice1, slice2 []int) []int {
	m := make(map[int]int)
	nn := make([]int, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

func main() {
	slice1 := []int{1, 2, 3, 6, 8}
	slice2 := []int{2, 3, 5, 6}

	in := intersect(slice1, slice2)
	fmt.Println("slice1与slice2的交集为：", in)

}
