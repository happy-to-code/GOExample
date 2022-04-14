package main

import (
	"fmt"
	"strings"
)

type LabelAdd struct {
	Creator int64    `json:"creator"` // 创建人(对应用户表的ID)
	Tags    []string `json:"tags"`    // 标签
}

func main() {
	// label := &LabelAdd{
	// 	Creator: 10,
	// 	Tags:    []string{"aa", "bb"},
	// }
	// fmt.Println(label)
	// fmt.Println("============================")
	// var list = []string{"11", "22", "33"}
	//
	// join := strings.Join(list, ",")
	// fmt.Println(join)

	var key1 = "/v1/2bsc2/casbin/get"
	// s := key1[4:]
	// fmt.Println(s)
	// index := strings.Index(s, "/")
	// fmt.Println(s[index:])
	var key2 = "/v1/*/casbin/get"
	fmt.Println(KeyMatch(key1, key2))

}

// KeyMatch key2 模板
func KeyMatch(key1, key2 string) bool {
	i := strings.Index(key2, "*")
	if i == -1 {
		return key1 == key2
	}
	split1 := strings.Split(key1, "/")
	split2 := strings.Split(key2, "/")
	if len(split2) != len(split1) {
		return false
	}
	for i := 1; i < len(split1); i++ {
		s2 := split2[i]
		s1 := split1[i]
		if s2 == "*" {
			continue
		}
		if s1 != s2 {
			return false
		}
	}
	return true
}
