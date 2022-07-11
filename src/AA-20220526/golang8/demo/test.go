package main

import "fmt"

// AddElem 1. 泛型的类型限制，在函数上直接申明该函数支持的多个类型
func AddElem[T int | string](params []T) (sum T) {
	for _, elem := range params {
		sum += elem
	}
	return
}

func main() {
	// 1. 在函数上声明泛型支持的多个类型
	// 1.1 传入支持的int
	intSum := AddElem([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Printf("测试1.1: 类型=%T，val=%+v\n", intSum, intSum)

	// 1.2 传入支持的string
	strSum := AddElem([]string{"静", "以", "修", "身", "，", "俭", "以", "养", "德"})
	fmt.Printf("测试1.2: 类型=%T，val=%+v", strSum, strSum)
}
