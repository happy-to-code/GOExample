package main

import (
	"fmt"
	"strings"
)

func main() {
	var list = []string{"A", "Bc", "dFFF"}

	fmt.Println(list)

	for i := range list {
		list[i] = strings.ToLower(list[i])
	}
	fmt.Println(list)
}
