package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math/big"
)

func main() {
	gasUsed1 := "0"
	gasLimit1 := "0000"
	// gasUsed
	gasUsed := new(big.Int)
	gasUsed.SetString(gasUsed1, 16)

	fmt.Println("gasUsed:", gasUsed)
	// gasLimit
	gasLimit := new(big.Int)
	gasLimit.SetString(gasLimit1, 16)

	fmt.Println("gasLimit:", gasLimit)
	// 计算gas使用百分比
	// 科学计算法
	decimal.DivisionPrecision = 2
	gas_used, _ := decimal.NewFromString(gasUsed.String())
	gas_limit, _ := decimal.NewFromString(gasLimit.String())

	fmt.Println("-->", gas_limit)
	fmt.Println("-->", gas_limit.String())
	fmt.Println("-->", gas_limit.String() == "0")
	percentageGasUsed := gas_used.Div(gas_limit).String()
	fmt.Println(percentageGasUsed)
}
