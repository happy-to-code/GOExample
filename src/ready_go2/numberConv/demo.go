package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	// uint64 18446744073709551615
	// 1510b3dc2b88a2b495711a61	6519377384383146429558889057
	// var a float64
	var s = "118d42d1e45dd210a15a0824"
	// var s = "1510b3d1"
	// newFromString, err := decimal.NewFromString(s)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("-->", newFromString)
	//
	// parseUint, err := strconv.ParseUint(s, 16, 32)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("===>", parseUint)

	// 118d42d1e45dd210a15a0824	5432019255541449266185242660
	// 118d42d1e45dd210a15a0824	5432019255541449864008488284
	// 118d42d1e45dd210a15a0824	5432019255541449266185242660
	ether := weiToEther(s)
	fmt.Println("===>", ether)
}

func weiToEther(hex string) *big.Float {
	var val = hex
	if strings.HasPrefix(hex, "0x") {
		val = hex[2:]
	}
	wei := new(big.Int)
	wei.SetString(val, 16)

	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)
	return f.Quo(fWei.SetInt(wei), big.NewFloat(Ether))
}

const (
	Wei   = 1
	GWei  = 1e9
	Ether = 1e18
)
