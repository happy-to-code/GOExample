package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"math/big"
	"strings"
)

func main() {
	// var a = 2
	// var b = 10
	//
	// fmt.Println(math.Pow10(a))
	// fmt.Println(math.Pow10(3))
	// fmt.Println(math.Pow10(4))
	// fmt.Println(math.Pow10(5))
	// fmt.Println(math.Pow10(6))
	// fmt.Println(math.Pow10(18))
	// fmt.Println(math.Pow10(1))
	// fmt.Println(math.Pow10(0))
	//
	// float := decimal.NewFromFloat(math.Pow10(18))
	// // 1000000000000000000
	// fmt.Println(float)

	var tokenBAmount = "183884303175509188243123141"
	var tokenB_total_amount = "5711310"

	tokenBAmount1, _ := decimal.NewFromString(tokenBAmount)
	tokenB_total_amount1, _ := decimal.NewFromString(tokenB_total_amount)
	s := tokenBAmount1.Div(decimal.NewFromFloat(math.Pow10(6))).DivRound(tokenB_total_amount1, 10).Mul(decimal.NewFromFloat(100)).String()
	fmt.Println(s)

}

func Int2String(v *big.Int, decimals int) string {
	str := v.Text(10)
	if decimals <= 0 {
		return str
	}
	if len(str) <= decimals {
		return strings.TrimRight(fmt.Sprintf("0.%s%s", strings.Repeat("0", decimals-len(str)), str), "0")
	}
	pos := len(str) - decimals
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%s.%s", str[:pos], str[pos:]), "0"), ".")
}

// subject:主体;account:账户;product:产品;transactionreport:交易报告;registration:登记;settlement:资金结算;infodisclosure:信批;finance:财务报表;regulation:监管;
