package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// computedAddr()
	nonce := uint64(6)
	contractAddr := crypto.CreateAddress(common.HexToAddress("0x36928500bc1dcd7af6a2b4008875cc336b927d57"), nonce)
	fmt.Printf("addr:%v\n", contractAddr.Hex())

	// 0xdac17f958d2ee523a2206206994597c13d831ec7
	// 0xdAC17F958D2ee523a2206206994597C13D831ec7
}

func computedAddr() {
	addr := common.HexToAddress("0x6ac7ea33f8831ea9dcc53393aaa88b25a785dbf0")

	nonce := uint64(135)
	contractAddr := crypto.CreateAddress(addr, nonce)
	fmt.Printf("addr:%v\n", contractAddr.Hex())
	//	tokenAddress:=common.HexToAddress("0xB66a603f4cFe17e3D27B87a8BfCaD319856518B8")
	createrAddress := common.HexToAddress("0x3482549fCa7511267C9Ef7089507c0F16eA1dcC1")
	contractAddr = crypto.CreateAddress(createrAddress, nonce)
	fmt.Printf("addr:%v\n", contractAddr.Hex())

}
