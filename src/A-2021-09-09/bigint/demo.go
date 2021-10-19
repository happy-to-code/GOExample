package main

import (
	"fmt"
	"strings"
)

func main() {
	// 140000000000000
	// 18446744073709551615
	var uintMaxStr = "18446744073709 551615"
	var amount = 18346744073709

	amountStr := fmt.Sprintf("%d%s", amount, "000000")
	fmt.Println("amountStr:", amountStr)

	if len(amountStr) > len(uintMaxStr) {
		fmt.Println("------max------")
		return
	}

	compare := strings.Compare(amountStr, uintMaxStr)
	if compare > 0 {
		fmt.Printf("mmmmm")
	} else {
		fmt.Println("lllllsdfsdfsdf")
	}

}

// 18446744073709551615
// 18346744073709000000
