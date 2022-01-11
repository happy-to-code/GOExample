package BLC

import (
	"fmt"
	"testing"
)

func TestCreateGenesisBlock(t *testing.T) {
	genesisBlock := CreateGenesisBlock("Hello World")
	fmt.Println(string(genesisBlock.PrevBlockHash))
	fmt.Println(genesisBlock.TimeStamp)
	fmt.Println(string(genesisBlock.Data))
	fmt.Println(string(genesisBlock.Hash))
	fmt.Println(genesisBlock.Height)
}
