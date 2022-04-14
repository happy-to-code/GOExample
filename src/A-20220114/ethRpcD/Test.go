package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {

	client, err := rpc.Dial("http://10.1.3.116:3334")
	if err != nil {
		fmt.Println("rpc.Dial err", err)
		return
	}

	var account []string
	err = client.Call(&account, "eth_accounts")
	fmt.Println("account:", account)
	var result interface{}
	// var result hexutil.Big
	err = client.Call(&result, "eth_getTransactionByBlockNumberAndIndex", "0x60fa8f", "0x1")
	// err = ec.c.CallContext(ctx, &result, "eth_getBalance", account, "latest")

	if err != nil {
		fmt.Println("client.Call err", err)
		return
	}
	fmt.Println("result:", result)
	// fmt.Printf("account[0]: %s\nbalance[0]: %s\n", account[0], result)
	// fmt.Printf("accounts: %s\n", account[0])
}
