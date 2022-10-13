package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	// s1 := "2828325396425257598899381"
	// s2 := "2828325396425257698899381"

	zero, _ := decimal.NewFromString("0")
	fmt.Println("zero:", zero)

	s1 := "0"
	s2 := "4"

	// s1Float, b := new(big.Float).SetString(s1)
	// if !b {
	// 	fmt.Println("s1转换失败")
	// }
	// s2Float, b := new(big.Float).SetString(s2)
	// if !b {
	// 	fmt.Println("s2转换失败")
	// }
	//
	// fmt.Println("-->", s1Float, "\t", s2Float)
	// sub := s2Float.Sub(s2Float, s1Float)
	// fmt.Println(sub)

	fromS1, err := decimal.NewFromString(s1)
	if err != nil {
		panic(err)
	}
	fromS2, err := decimal.NewFromString(s2)
	if err != nil {
		panic(err)
	}
	fmt.Println("==>", fromS1, "\t", fromS2)
	sub2 := fromS2.Sub(fromS1)
	fmt.Println(sub2)

	fmt.Println("==============>", fromS1)
	fmt.Println("==============>", fromS1 == zero)
	mul := sub2.DivRound(fromS1, 6).Mul(decimal.NewFromFloat(100))
	fmt.Println(mul)
}
