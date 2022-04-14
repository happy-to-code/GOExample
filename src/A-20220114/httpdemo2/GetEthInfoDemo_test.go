package main

import (
	"fmt"
	"testing"
)

func Test_updateAccountAddress(t *testing.T) {
	type args struct {
		from  string
		to    string
		stamp int64
		i     int
		index string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				from:  "from1",
				to:    "to1",
				stamp: 1539402630,
				i:     456,
				index: "0x2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateAccountAddress("from1", "to1", 1539402630, 456, "0x2")
		})
	}
	fmt.Println("----------")
}

func Test_selectAssetBalanceByAddress(t *testing.T) {
	err, balance := selectAssetBalanceByAddress("123ss1")
	if err != nil {
		panic(err)
	}
	fmt.Println("==>", balance)
	fmt.Println("==>", balance.LastTradeTxHash[0])
	fmt.Println("==>", balance.LastTradeTxHash[1])
}

func Test_selectMaxHasHandledBlockNum(t *testing.T) {
	balance := selectMaxHasHandledBlockNum()
	if err != nil {
		panic(err)
	}
	fmt.Println("==>", balance)
}
