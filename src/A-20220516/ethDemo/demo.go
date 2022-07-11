package main

const (
	keyStorePath = `your keyStorePath`
	fromKeyStore = `your from keyStore`
	toAddress    = "0x477257735cCEF9C7d42856649c7149865D04eDeb"
	chainId      = 7851254
	chtAddress   = "0xb9bD7405797CfFBc7e57309444b4af89c39cA92c" // 合约地址
)

// func main() {
// 	// 服务器地址
// 	conn, err := ethclient.Dial("http://172.16.30.174:8545")
// 	if err != nil {
// 		fmt.Println("Dial err", err)
// 		return
// 	}
// 	defer conn.Close()
// 	// 创建合约对象
// 	chtToken, err := NewChttoken(common.HexToAddress(chtAddress), conn)
// 	if err != nil {
// 		fmt.Println("newChttoken error", err)
// 	}
// 	// 调用查询方法
// 	// symbol, err := chtToken.Decimals(nil)
// 	// if err != nil {
// 	//    fmt.Println("invoke error", err)
// 	// }
// 	// fmt.Println(symbol)
//
// 	res1, err := chtToken.BalanceOf(&bind.CallOpts{
// 		Pending:     false,
// 		From:        common.Address{},
// 		BlockNumber: nil,
// 		Context:     nil,
// 	}, common.HexToAddress(toAddress))
// 	if err != nil {
// 		fmt.Println("BalanceOf error", err)
// 	}
// 	fmt.Println(res1)
//
// 	// 调用其他方法，需要组装from
// 	// 解锁对应账户
// 	fromKey, err := keystore.DecryptKey([]byte(fromKeyStore), "abc123")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	auth, err := bind.NewKeyedTransactorWithChainID(fromKey.PrivateKey, new(big.Int).SetInt64(chainId))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// fromPrivateKey := fromKey.PrivateKey
// 	// fromPublicKey := fromPrivateKey.PublicKey
// 	// fromAddr := crypto.PubkeyToAddress(fromPublicKey)
//
// 	tx, err := chtToken.Transfer(&bind.TransactOpts{
// 		From: auth.From,
// 		// Nonce:     nil,
// 		Signer: auth.Signer,
// 		// Value:     nil,
// 		// GasPrice:  nil,
// 		// GasFeeCap: nil,
// 		// GasTipCap: nil,
// 		// GasLimit:  0,
// 		// Context:   nil,
// 		// NoSend:    false,
// 	}, common.HexToAddress(toAddress), new(big.Int).SetInt64(1e18))
//
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(tx.Hash())
// 	// 等待挖矿完成
// 	receipt, err := bind.WaitMined(context.Background(), conn, tx)
// 	if err != nil {
// 		fmt.Println("WaitMined error", err)
// 	}
// 	fmt.Println(receipt.BlockNumber)
//
// 	// 转账之后查询余额
// 	res2, err := chtToken.BalanceOf(&bind.CallOpts{
// 		Pending:     false,
// 		From:        common.Address{},
// 		BlockNumber: nil,
// 		Context:     nil,
// 	}, common.HexToAddress(toAddress))
// 	if err != nil {
// 		fmt.Println("BalanceOf error", err)
// 	}
// 	fmt.Println(res2)
// }
