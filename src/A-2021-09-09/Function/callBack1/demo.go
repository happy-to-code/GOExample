package main

import (
	"fmt"
	"time"
)

type Callback func(num int) int

// 提供一个接口，让外部去实现
func test(x, y int, callback Callback) {
	fmt.Println(x, "--->", y, "保存数据库……")
	time.Sleep(time.Second * 2)

	i := callback(y)
	fmt.Println("i===>", i)
}

func main() {
	test(10, 1, count)

}

func count(a int) int {
	fmt.Println("count:", a)
	return a + 99
}
