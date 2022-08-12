package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "123"
	if strings.Contains(s, "1") {
		fmt.Println("1")
	} else if strings.Contains(s, "2") {
		fmt.Println("2")
	} else if strings.Contains(s, "3") {
		fmt.Println("3")
	} else {
		fmt.Println(4)
	}
}
