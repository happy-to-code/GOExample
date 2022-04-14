package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
)

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
func main() {
	fmt.Println(weiToEther("0x40200048201"))

	// 0.000004406636741121
	// 4.406636741121e-06
}
