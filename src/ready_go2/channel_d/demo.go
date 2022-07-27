package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValue(values chan int) {
	// 设置随机种子,避免随机函数生成相同的值
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)

	fmt.Printf("计算随机值: {%d}\n", value)

	// 往通道发送值
	values <- value
}

func main() {
	// 创建int类型通道，只能传入int类型值
	values := make(chan int)
	defer close(values)

	go CalculateValue(values)

	// 从通道接收值
	value := <-values
	fmt.Println(value)
}
