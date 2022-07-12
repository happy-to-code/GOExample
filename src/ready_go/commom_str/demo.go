package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abcdef"
	fmt.Println(strings.Index(str, "abcdr"))

	strs := []string{"flower", "flow", "flight"}
	prefix := longestCommonPrefix(strs)
	fmt.Println(prefix)
}

// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串""。
// 示例 1：
// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
// 示例 2：
//
// 输入：strs = ["dog","racecar","car"]
// 输出：""
// 解释：输入不存在公共前缀。
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var target = strs[0]
	for _, str := range strs {
		if str == "" {
			return ""
		}

		for !strings.HasPrefix(str, target) {
			target = target[:len(target)-1]
		}
		if target == "" {
			return ""
		}
	}

	return target
}
