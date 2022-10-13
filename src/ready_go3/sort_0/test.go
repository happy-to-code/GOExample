package main

import "fmt"

func main() {
	var nums = []int{5, 1, 2, 4, 3}
	fmt.Println("BubbleSort:\t", BubbleSort(nums))
	fmt.Println("SelectionSort:\t", SelectionSort(nums))
	fmt.Println("InsertSort:\t", InsertSort(nums))
	fmt.Println("ShellSort:\t", ShellSort(nums))

}

// BubbleSort 首先从数组的第一个元素开始到数组最后一个元素为止，
// 对数组中相邻的两个元素进行比较，如果位于数组左端的元素大于数组右端的元素，则交换这两个元素在数组中的位置。
func BubbleSort(nums []int) []int {
	length := len(nums)
	if length == 0 || length == 1 {
		return nums
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

// SelectionSort 每一趟在n-i+1(i=1,2,...,n-1)个记录中选取关键字最小的记录作为有序序列中第i个记录。
func SelectionSort(nums []int) []int {
	length := len(nums)
	if length == 0 || length == 1 {
		return nums
	}
	for i := 0; i < length; i++ {
		index := i
		for j := i + 1; j < length; j++ {
			if nums[j] < nums[index] {
				index = j
			}
		}
		if index != i {
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
	return nums
}

// InsertSort 插入排序的基本思想就是将无序序列插入到有序序列中。
// 例如要将数组arr=[4,2,8,0,5,1]排序，可以将4看做是一个有序序列(图中用蓝色标出)，
// 将[2,8,0,5,1]看做一个无序序列。无序序列中2比4小，于是将2插入到4的左边，此时有序序列变成了[2,4]，
// 无序序列变成了[8,0,5,1]。无序序列中8比4大，于是将8插入到4的右边，有序序列变成了[2,4,8],
// 无序序列变成了[0,5,1]。以此类推，最终数组按照从小到大排序
func InsertSort(nums []int) []int {
	length := len(nums)
	if length == 0 || length == 1 {
		return nums
	}

	for i := 1; i < length; i++ {
		var j int
		if nums[i] < nums[i-1] {
			temp := nums[i]
			for j := i - 1; j >= 0 && temp < nums[j]; j-- {
				nums[j+1] = nums[j]
			}
			nums[j+1] = temp
		}
	}
	return nums
}

// ShellSort 希尔排序(Shell's Sort)在插入排序算法的基础上进行了改进，
// 算法的时间复杂度与前面几种算法相比有较大的改进。
// 其算法的基本思想是：先将待排记录序列分割成为若干子序列分别进行插入排序，待整个序列中的记录"基本有序"时，再对全体记录进行一次直接插入排序。
func ShellSort(nums []int) []int {
	length := len(nums)
	if length == 0 || length == 1 {
		return nums
	}
	var k int
	increasement := length
	for increasement > 1 {
		increasement = increasement/3 + 1 // 分组数量
		for i := 0; i < increasement; i++ {
			for j := i + increasement; j < length; j += increasement {
				if nums[j] < nums[j-increasement] {
					temp := nums[j]
					for k := j - increasement; k >= 0 && temp < nums[k]; k -= increasement {
						nums[k+increasement] = nums[k]
					}
					nums[k+increasement] = temp
				}
			}
		}
	}

	return nums
}
