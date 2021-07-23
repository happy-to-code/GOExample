package main

import (
	"fmt"
	"strings"
)

func main() {
	// s := "2020年6 月"
	// s2 := strings.ReplaceAll(s, "年", "")
	// print(s2)

	ss := " ss dsd  "
	if ss = strings.TrimSpace(ss); ss == " ss dsd" {
		return
	}
	fmt.Println("--success--")
}
