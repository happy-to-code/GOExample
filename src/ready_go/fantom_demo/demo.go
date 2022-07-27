package main

import (
	"context"
	"fmt"
	"github.com/Fantom-foundation/go-ethereum/ethclient"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("http://10.1.9.222:18545")
	if err != nil {
		panic(err)
	}
	blockByNumber, err := client.BlockByNumber(context.Background(), big.NewInt(30))
	if err != nil {
		panic(err)
	}
	fmt.Printf("blockByNumber:%+v\n", blockByNumber)
	fmt.Printf("blockHash:%x\n", blockByNumber.Hash())
}
