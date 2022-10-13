package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5, 6, 7}
	traverse(a, 0)

}

/* 递归遍历数组 */
func traverse(arr []int, i int) {
	if i == len(arr) {
		return
	}

	fmt.Println(arr[i])
	// 前序位置
	traverse(arr, i+1)
	// 后序位置
	// fmt.Println(arr[i])

}
