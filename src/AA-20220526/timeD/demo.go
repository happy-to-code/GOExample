package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	time := time.Now().Add(-time.Second * 300)
	fmt.Println(time)
	format := time.Format("2006-01-02 15:04:05")
	fmt.Println(`format: ${format}`, format)

	idcard := "33192318901205636X"
	fmt.Println(CheckIdCard(idcard))
}

// CheckIdCard 检验身份证
func CheckIdCard(card string) bool {
	regRuler := "(^[1-9]\\d{5}(18|19|20|(3\\d))\\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\\d{3}[0-9Xx]$)"
	// 正则调用规则
	reg := regexp.MustCompile(regRuler)
	// 返回 MatchString 是否匹配
	return reg.MatchString(card)
}
