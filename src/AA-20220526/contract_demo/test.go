package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 合约地址 0x7479f75b06A4606721a35F115B28a47b99F99931
func main() {
	host := "http://127.0.0.1:8545"
	conn, err := ethclient.Dial(host)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	contractAddr := "0x7479f75b06A4606721a35F115B28a47b99F99931"
	msgSenderDemoIns, err := NewMsgSenderDemo(common.HexToAddress(contractAddr), conn)
	if err != nil {
		fmt.Println("NewMsgSenderDemo,err:", err.Error())
		panic(err)
	}
	sender, err := msgSenderDemoIns.GetSender(nil)
	if err != nil {
		fmt.Println("msgSenderDemoIns:", err)
		panic(err)
	}
	fmt.Println("sender:", sender)
	fmt.Println("==================================================================")
	// sender2, err := msgSenderDemoIns.SetSender2(nil)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("sender2:", sender2)
	//
	// sender33, err := msgSenderDemoIns.GetSender(nil)
	// if err != nil {
	// 	fmt.Println("msgSenderDemoIns333:", err)
	// 	panic(err)
	// }
	// fmt.Println("sender33:", sender33)

}
