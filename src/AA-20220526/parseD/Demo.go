package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s = "0.03826412345678911121314"
	float, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(float)

	var ss = " >= 0.03 "
	fmt.Printf("--->%s<-----\n", ss)
	// ss = strings.TrimSpace(ss)
	ss = strings.ReplaceAll(ss, " ", "")
	fmt.Printf("--->%s<-----\n", ss)
	fmt.Println(strings.Contains(ss, ">="))
	fmt.Println(strings.Contains(ss, ">"))
}
