package main

import (
	"fmt"
	"math/big"
)

type TxHistory struct {
	Txhash         string     `1.json:"txhash"`
	From           string     `1.json:"from"`
	To             string     `1.json:"to"`
	Amount         *big.Float `1.json:"amount"`
	OriginalAmount string     `1.json:"original_amount"`
	Token          string     `1.json:"token"`
	GasUsed        int64      `1.json:"gas_used"`
	GasPrice       int64      `1.json:"gas_price"`
	ChainName      string     `1.json:"chain_name"`
	CreatedTime    int64      `1.json:"created_time"`
	BlockNumber    int64      `1.json:"block_number"`
	Status         int        `1.json:"status"`
}

func main() {
	sql := "INSERT INTO \"tx_history\"(\"txhash\", \"from\", \"to\", \"amount\", \"token\", \"gas_used\", \"gas_price\", \"chain_name\", \"created_time\", \"block_number\", \"status\") VALUES "

	txHistorys := []TxHistory{
		{
			Txhash:         "sdfds",
			From:           "dfsdgtg",
			To:             "gfgfgg",
			Amount:         new(big.Float).SetFloat64(12),
			OriginalAmount: "gfgfgdgsr",
			Token:          "er",
			GasUsed:        23,
			GasPrice:       56,
			ChainName:      "eth",
			CreatedTime:    123,
			BlockNumber:    2,
			Status:         1,
		},
		{
			Txhash:         "sdf121ds",
			From:           "dfs23232dgtg",
			To:             "gf323232gfgg",
			Amount:         new(big.Float).SetFloat64(568),
			OriginalAmount: "gfgfgere33dgsr",
			Token:          "er343",
			GasUsed:        231,
			GasPrice:       5886,
			ChainName:      "eth",
			CreatedTime:    123435345,
			BlockNumber:    3,
			Status:         0,
		},
	}
	var s = ""
	for i, history := range txHistorys {
		sprintf := fmt.Sprintf("(%s,%s,%s,%f,%s,%d,%d,%s,%d,%d,%d)", history.Txhash, history.From, history.To, history.Amount, history.Token, history.GasUsed, history.GasPrice, history.ChainName, history.CreatedTime, history.BlockNumber, history.Status)

		if i == 0 {
			s = fmt.Sprintf("%s%s", s, sprintf)
		} else {
			s = fmt.Sprintf("%s,%s", s, sprintf)
		}
	}
	sql = sql + s

	fmt.Println(sql)
}
