package main

import (
	"fmt"
	"strings"
)

type ProcessBasename func(string) string

// basename函数 移除字符串的路径部分和.后缀 只取文件名
func basename(str string) string {
	// 去除路径部分
	slash := strings.LastIndex(str, "/") // 如果没有找到"/" LastIndex返回-1
	str = str[slash+1:]
	// 取出.的前一部分
	if dot := strings.LastIndex(str, "."); dot > 0 {
		str = str[:dot]
	}
	return str
}

// GetBaseName 这个函数使用了函数变量作为参数
func GetBaseName(str string, myBasenameFunc func(string) string) string {
	return myBasenameFunc(str)
}

// GetBaseName1 这个函数使用了type关键字自定义的类型(底层是一种函数类型)
func GetBaseName1(str string, processBasename ProcessBasename) string {
	return processBasename(str)
}

// 函数变量被当做函数参数传递
func main() {
	str := "a/b/c.go"
	// 将函数basename作为实参传递给GetBaseName
	// 要注意，函数变量basename的函数签名要和形参myBasenameFunc的函数签名完全一致
	// 在这里，函数签名是指 函数的参数和返回值
	fmt.Println(GetBaseName(str, basename))
	str1 := "c.d.go"
	fmt.Println(GetBaseName1(str1, basename))
}
