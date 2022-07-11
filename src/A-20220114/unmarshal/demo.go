package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var s = `{
"blockHash": "0xc6dbd3d711ed918da3fa83c7ebf99cfe2b4ebdf5ecb5356b67c21f05c9f0bda9",
"blockNumber": "0x634291",
"from": "0xea674fdde714fd979de3edf0f56aa9716b898ec8",
"gas": "0xc350",
"gasPrice": "0x3b9aca00",
"hash": "0x3f6de45212e26beca7a9711935dc722f6b6617095d085e58eb5f84f9a3ee3f90",
"input": "0x",
"nonce": "0xedda92",
"to": null,
"transactionIndex": "0x0",
"value": "0xb1b3fe3f92b2c4",
"type": "0x0",
"v": "0x25",
"r": "0x137e7c666f53455078bfa66f01747693fabaaef54078e44d8b561cdee0115a9c",
"s": "0x2f4c6f91a1eb9f122fd3f219cecd205c02c4295eaa1736e2cb1a4bf8dad623e"
}`

	var t Transaction
	err := json.Unmarshal([]byte(s), &t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", t)
	fmt.Println("ssdf:", t.To)

	sprintf := fmt.Sprintf("(%d,%s)", 12, "xm")
	fmt.Println(sprintf)
}

type Transaction struct {
	BlockHash        string `1.json:"blockHash"`
	BlockNumber      string `1.json:"blockNumber"`
	From             string `1.json:"from"`
	Gas              string `1.json:"gas"`
	GasPrice         string `1.json:"gasPrice"`
	Hash             string `1.json:"hash"`
	Input            string `1.json:"input"`
	Nonce            string `1.json:"nonce"`
	To               string `1.json:"to"`
	TransactionIndex string `1.json:"transactionIndex"`
	Value            string `1.json:"value"`
	Type             string `1.json:"type"`
	V                string `1.json:"v"`
	R                string `1.json:"r"`
	S                string `1.json:"s"`
}
