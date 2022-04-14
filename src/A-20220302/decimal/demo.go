package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	initAmount, _ := decimal.NewFromString("0")
	a, _ := decimal.NewFromString("0.1")
	b, _ := decimal.NewFromString("0.35")
	c, _ := decimal.NewFromString("1")

	s := initAmount.Add(a).Add(b).Add(c).String()
	fmt.Println(s)

}
