package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	var decimals = 18
	v, _ := big.NewInt(0).SetString("2412dd1a4e600", 16)

	str := v.Text(10)

	if len(str) <= decimals {
		fmt.Println("1===>", strings.TrimRight(fmt.Sprintf("0.%s%s", strings.Repeat("0", decimals-len(str)), str), "0"))
	}
	pos := len(str) - decimals
	fmt.Println("2===>", strings.TrimRight(strings.TrimRight(fmt.Sprintf("%s.%s", str[:pos], str[pos:]), "0"), "."))
}
