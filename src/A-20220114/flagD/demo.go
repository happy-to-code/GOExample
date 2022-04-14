package main

import (
	"flag"
	"fmt"
)

func main() {
	// 定义接受参数的变量
	var start int64
	var end int64

	flag.Int64Var(&start, "s", start, "开始区块")
	flag.Int64Var(&end, "e", end, "结束区块")

	// 解析命令行参数写入注册的flag里,在所有flag都注册之后，这一步一定不能少
	flag.Parse()

	fmt.Println("s", start+1)
	fmt.Println("e", end+1)

	for i := 0; i < 1; i++ {
		fmt.Println("-----")
		fmt.Println(i)
	}
}
