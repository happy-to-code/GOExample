package main

import "fmt"

func main() {
	var nums = []int{1, 1, 2, 2, 3, 4}
	dupl := removeDupl(nums)
	fmt.Println(dupl)
}

func removeDupl(nums []int) (res []int) {
	if len(nums) == 0 {
		return
	}
	var slow = 0
	var fast = 0
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	res = nums[:slow+1]

	return
}
