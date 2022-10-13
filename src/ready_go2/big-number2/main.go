package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {

	zero, _ := decimal.NewFromString("0")
	fmt.Println("zero:", zero)
	zero2 := decimal.NewFromFloat(0)

	s1 := "0"
	// s2 := "4"

	fromS1, err := decimal.NewFromString(s1)
	if err != nil {
		panic(err)
	}

	fmt.Println("==============>", fromS1)
	fmt.Println("==============>", fromS1 == zero)
	fmt.Println("==============>", fromS1.Equal(zero))
	fmt.Println("==============>", fromS1.Equal(zero2))

}
