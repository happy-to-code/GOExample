package main

import "fmt"

func main() {
	arr := []string{}
	arr = append(arr, "main")
	fmt.Printf("1 %p %p\n", &arr, arr)
	// modifyFunc1(&arr)
	// fmt.Println("slice：main函数：", arr)

	modifyFunc2(&arr)
	fmt.Println("slice：main函数2：", arr)
}

// 值未改变
func modifyFunc1(arrParam *[]string) {
	fmt.Printf("2 %p %p\n", &arrParam, arrParam)
	tmp := []string{"tmp modifyFunc1"}
	arrParam = &tmp
	fmt.Printf("3 %p %p\n", &arrParam, arrParam)
	fmt.Println("slice：modifyFunc1函数：", arrParam)
}

func modifyFunc2(arrParam *[]string) {
	fmt.Printf("2 %p %p\n", &arrParam, arrParam)
	tmp := []string{"tmp modifyFunc2"}
	*arrParam = tmp
	fmt.Printf("3 %p %p\n", &arrParam, arrParam)
	fmt.Println("slice：modifyFunc1函数：", arrParam)
}
