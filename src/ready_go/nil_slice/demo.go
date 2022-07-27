package main

import "fmt"

func main() {
	var list []string

	list = make([]string, 2)
	list[0] = "hello"
	list[1] = "world"
	fmt.Println(list)

	var s = "你好abc啊哈哈"
	fmt.Println(string(reverse([]rune(s))))
}
func reverse(s []rune) []rune {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
