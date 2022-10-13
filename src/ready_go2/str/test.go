package main

import "fmt"

func main() {
	var nums = []int64{1, 4, 6, 90, 34}
	var blockNumberStr string
	for i, num := range nums {
		var temp string
		if i == 0 {
			temp = fmt.Sprintf("%d", num)
		} else {
			temp = fmt.Sprintf("%s%d", ",", num)
		}
		blockNumberStr = blockNumberStr + temp
	}

	fmt.Println(blockNumberStr)
}
