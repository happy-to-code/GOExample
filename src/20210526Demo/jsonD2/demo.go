package main

import (
	"encoding/json"
	"fmt"
)

type History struct {
	TxId       string // 区块交易ID
	Account    string // 账号，即地址信息
	PointType  string // 积分信息(平台积分:platform   贡献积分:contribute)
	TxType     string // 交易类型(增加积分:credit     扣减积分:debit)
	Amount     int64  // 积分发生额度
	Balance    int64  // 积分变动后余额(当时的余额)
	ReasonCode string // 触发积分业务代码
	TxTime     int64  // 交易时间
}

func main() {
	var h = History{
		TxId:       "1234",
		Account:    "abc",
		PointType:  "",
		TxType:     "credit",
		Amount:     10,
		Balance:    10,
		ReasonCode: "test",
		TxTime:     1624425609,
	}
	marshal, _ := json.Marshal(h)
	fmt.Println(string(marshal))
	// s := "{\"TxId\":\"1234\",\"Account\":\"abc\",\"PointType\":\"\",\"TxType\":\"credit\",\"Amount\":10,\"Balance\":10,\"ReasonCode\":\"test\",\"TxTime\":1624425609}"
	uniqueKey := "unique_id_Z000351"
	fmt.Println("len:", len([]byte(uniqueKey)))
}
