package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(binaryQuery(nums, 5))
}

func binaryQuery(nums []int, target int) (a int) {
	a = -1
	var left = 0
	var right = len(nums) - 1

	for left <= right {
		// var mid = (left + right) / 2
		var mid = left + (right-left)/2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			a = mid
			return
		}
	}

	return
}
