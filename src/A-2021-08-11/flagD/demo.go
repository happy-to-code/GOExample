package main

import (
	"flag"
	"fmt"
)

// 解析命令行参数
func main() {
	// 定义接受参数的变量
	var name string
	var password string
	var host string
	var port int

	// 使用flag函数接受参数
	/**
	参数一 接收的变量名 引用类型
	参数二 指定命令行用什么接收 string
	参数三 默认值，注意根据变量类型设置默认值 type
	参数四 说明，string
	*/
	flag.StringVar(&name, "u", "", "用户名")
	flag.StringVar(&password, "p", "", "用户密码")
	flag.StringVar(&host, "h", "localhost", "主机名")
	flag.IntVar(&port, "a", 3306, "主机端口")

	// 解析命令行参数写入注册的flag里,在所有flag都注册之后，这一步一定不能少
	flag.Parse()

	// 参看结果
	fmt.Printf("name:%v\n", name)
	fmt.Printf("password:%v\n", password)
	fmt.Printf("host:%v\n", host)
	fmt.Printf("port:%v\n", port)

	// sysType := runtime.GOOS

}
