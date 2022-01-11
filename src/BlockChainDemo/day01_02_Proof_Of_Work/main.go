package main

import (
	"BlockChainDemo/day01_02_Proof_Of_Work/BLC"
	"fmt"
)

func main() {
	// 4.测试添加新区块
	blockChain := BLC.CreateBlockChainWithGenesisBlock("Genesis Block..")
	blockChain.AddBlockToBlockChain("Send 100RMB To Wangergou", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	blockChain.AddBlockToBlockChain("Send 300RMB To lixiaohua", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	blockChain.AddBlockToBlockChain("Send 500RMB To rose", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	fmt.Println("blockChain:", blockChain)
	pow := BLC.NewProofOfWork(blockChain.Blocks[1])
	fmt.Printf("%v\n", pow.IsValid())

}
