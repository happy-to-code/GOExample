package main

import (
	"fmt"
	"strings"
)

type Boy struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	txTime := 1651892056090
	txTimeStr := fmt.Sprintf("%d", txTime)
	if len(txTimeStr) == 13 {
		txTime = txTime / 1000
	}
	fmt.Println("txtime", txTime)

	bys := []Boy{
		{"XM", 12},
		{"XH", 18},
		{"ZS", 15},
	}
	fmt.Println(bys)
	for i, by := range bys {
		by.Age = by.Age + 1
		bys[i] = by
	}
	fmt.Println(bys)
	var list = []string{"钱***", "钱**", "弘艺丰艺术品有限公司", "李公司", "张三", "刘岳魏", "司徒雷登", "王五万五李四"}
	for _, s := range list {
		info := hiddenKeyInfo(s)
		fmt.Println(s, "-->", info)
	}

	// var t = 1657174123010
	fmt.Println(len("钱**"))
}
func hiddenKeyInfo(secret string) string {
	length := len(secret)
	if (length > 9 && strings.Contains(secret, "公司")) || length <= 3 || strings.Contains(secret, "*") { // 超过3个汉字或者1个字
		return secret
	}

	// 2个汉字
	if length == 6 {
		familyName := secret[:3]
		return familyName + "*"
	}
	// 3个汉字 或以上
	if length >= 9 {
		familyName := secret[:3]
		lastSingleName := secret[length-3:]
		return familyName + "*" + lastSingleName
	}

	return secret
}
