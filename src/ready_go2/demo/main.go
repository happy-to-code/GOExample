package main

import "fmt"

func main() {
	var list = []int{1, 0, 2, 5, 0}

	var j = 0
	for i := 0; i < len(list); i++ {
		if i == j && list[i] == 0 {
			continue
		}
		if list[i] != 0 {
			list[j], list[i] = list[i], list[j]
			j++
		}
	}
	fmt.Println(list)
}
