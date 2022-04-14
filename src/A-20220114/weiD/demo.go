package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"github.com/shopspring/decimal"
	"math/big"
)

func main() {
	// var s = "0x3ad7818917a8000"
	// ss := weiToEther(s)
	// fmt.Println("1====", ss)
	// fmt.Printf("2====%d\n", ss)
	// fmt.Println("3====", ss.String())
	// fmt.Println("-----------------------------")

	// 0.00671196"
	// m := new(big.Int).SetInt64(63209)
	// n := new(big.Int).SetInt64(82802353893)
	// mul := m.Mul(m, n)
	// fmt.Println(mul.String())
	//
	// gasPrice := new(big.Int)
	// setString, _ := gasPrice.SetString("1000000000000000000",10)
	// fmt.Println(setString)

	// s := "0x"
	// s := "0x214e8348c4f0000"
	s := "0xda6283"
	ether := weiToEther(s)
	fmt.Println(ether)
	fmt.Println(ether.String())

	// 1.4312067e-11
	// 0.000000000014312067
	// 0.000000000014312067
	decimal.DivisionPrecision = 18
	fromString, _ := decimal.NewFromString(ether.String())
	fmt.Println(fromString)

}
func weiToEther(hex string) *big.Float {
	val := hex[2:]
	wei := new(big.Int)
	wei.SetString(val, 16)

	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)
	return f.Quo(fWei.SetInt(wei), big.NewFloat(params.Ether))
}
