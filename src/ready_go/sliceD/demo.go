package main

import "fmt"

func main() {
	test1()
	fmt.Println("===============================")
	test2()

}

func test1() {
	aList := [3]int{1, 2, 3}
	for k, v := range aList {
		if k == 0 {
			aList[0], aList[1] = 100, 200
			fmt.Println("数组1：", aList)
		}
		aList[k] = v + 100
	}
	fmt.Println("数组2", aList)
}

func test2() {
	aList := []int{1, 2, 3}
	for k, v := range aList {
		if k == 0 {
			aList[0], aList[1] = 100, 200
			fmt.Println("切片1：", aList)
		}
		aList[k] = v + 100
	}
	fmt.Println("切片2", aList)
}
