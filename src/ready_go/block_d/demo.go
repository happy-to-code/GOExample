package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	"time"
)

func main() {
	host := "http://10.1.9.222:18545"
	client, err := ethclient.Dial(host)
	if err != nil {
		panic(err)
	}
	blockByNumber, err := client.BlockByNumber(context.Background(), big.NewInt(30))
	if err != nil {
		panic(err)
	}
	fmt.Printf("blockByNumber:%+v\n", blockByNumber)
	info := makeBlockInfo(blockByNumber)
	fmt.Printf("info:%+v\n", info)
	fmt.Printf("===>blockHash:%x,%+v\n", blockByNumber.Hash(), blockByNumber.Transactions())
}

var (
	TimeZone_Beijing, _ = time.LoadLocation("Asia/Shanghai")
)

func makeBlockInfo(block *types.Block) BlockMode {
	return BlockMode{
		BlockNumber: block.Number().Int64(),
		Hash:        strings.ToLower(block.Hash().Hex()),
		ParentHash:  strings.ToLower(block.ParentHash().Hex()),
		Miner:       strings.ToLower(block.Coinbase().Hex()),
		Difficulty:  block.Difficulty().Text(16),
		GasLimit:    block.GasLimit(),
		GasUsed:     block.GasUsed(),
		CreatedTime: time.Unix(int64(block.Time()), 0).In(TimeZone_Beijing).Format("2006-01-02 15:04:05"),
		CreatedAt:   time.Now(),
		Timestamp:   int64(block.Time()),
		TxCount:     block.Transactions().Len(),
		Nonce:       fmt.Sprintf("%d", block.Nonce()),
	}
}

type BlockMode struct {
	tableName   struct{}  `pg:"block_info"`
	BlockNumber int64     `pg:"block_number"`
	Hash        string    `pg:"block_hash"`
	ParentHash  string    `pg:"parent_hash"`
	Miner       string    `pg:"miner"`
	Difficulty  string    `pg:"difficulty"`
	GasLimit    uint64    `pg:"gas_limit"`
	GasUsed     uint64    `pg:"gas_used"`
	CreatedTime string    `pg:"created_time"`
	Timestamp   int64     `pg:"timestamp"`
	Nonce       string    `pg:"nonce"`
	TxCount     int       `pg:"tx_count"`
	CreatedAt   time.Time `pg:"created_at"`
}
