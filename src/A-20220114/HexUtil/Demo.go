package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
)

func weiToEther(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)
	return f.Quo(fWei.SetInt(wei), big.NewFloat(params.Ether))
}
func main() {
	// 899989998999899989999000
	// 899989998999899989999000
	// 1539404222
	// var s = "5bc171be"
	var s = "12a05f200"
	// 1,000,000,000,000,000,000
	fmt.Println("s:", s)

	wei := new(big.Int)
	wei.SetString(s, 16)
	fmt.Println("wei:", wei)
	fmt.Println("wei:", wei.Int64())
	ether := weiToEther(wei)
	fmt.Println("--", ether)
	f, accuracy := ether.Float64()
	fmt.Println(f, "==", accuracy)
	// ether.

	// parseInt, err := strconv.ParseInt(s, 16,64)
	// if err != nil {
	// 	panic(err)
	// }
	// // n2 := uint32(parseInt)
	// fmt.Println("10进制：", parseInt)

	// hex := "0xC40C5253"
	// val := hex[2:]
	//
	// n, err := strconv.ParseUint(val, 16, 32)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// n2 := uint32(n)
	// fmt.Print(n2)
}
