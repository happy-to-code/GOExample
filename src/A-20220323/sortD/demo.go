package main

import (
	"fmt"
	"sort"
)

type TradeInfo struct {
	ArtworkNo         string `json:"artworkNo"`         // 艺术品编号
	TransactionDate   int64  `json:"transactionDate"`   // 交易日期   十位：1651804856
	TransactionAmount string `json:"transactionAmount"` // 交易金额
	Seller            string `json:"seller"`            // 卖方姓名
	Buyer             string `json:"buyer"`             // 买方姓名
	AssuranceUnit     string `json:"assuranceUnit"`     // 鉴证单位
	DepositTime       int64  `json:"depositTime"`       // 存证时间
	OrderNo           string `json:"orderNo"`           // 订单编号
}

func main() {
	var trades = []TradeInfo{
		{
			ArtworkNo:         "no458988",
			TransactionDate:   1551004856,
			TransactionAmount: "82000万",
			Seller:            "王五",
			Buyer:             "钱**",
			AssuranceUnit:     "弘艺丰",
			DepositTime:       1551004856,
			OrderNo:           "no4596609",
		},
		{
			ArtworkNo:         "no458988",
			TransactionDate:   1551004856,
			TransactionAmount: "82000万",
			Seller:            "王五",
			Buyer:             "钱**",
			AssuranceUnit:     "弘艺丰",
			DepositTime:       1551004856,
			OrderNo:           "no4596609",
		},
		{
			ArtworkNo:         "no458988",
			TransactionDate:   1551004856,
			TransactionAmount: "82000万",
			Seller:            "王五",
			Buyer:             "钱**",
			AssuranceUnit:     "弘艺丰",
			DepositTime:       1551004856,
			OrderNo:           "no4596609",
		},
		{
			ArtworkNo:         "no458988",
			TransactionDate:   1551004856,
			TransactionAmount: "82000万",
			Seller:            "王五",
			Buyer:             "钱**",
			AssuranceUnit:     "弘艺丰",
			DepositTime:       1551004856,
			OrderNo:           "no4596609",
		},
		{
			ArtworkNo:         "no458990",
			TransactionDate:   1551004856,
			TransactionAmount: "82000万",
			Seller:            "王五",
			Buyer:             "钱**",
			AssuranceUnit:     "弘艺丰",
			DepositTime:       1551004856,
			OrderNo:           "no4596609",
		},
		{
			ArtworkNo:         "no458990",
			TransactionDate:   1551004856,
			TransactionAmount: "82000万",
			Seller:            "王五",
			Buyer:             "钱**",
			AssuranceUnit:     "弘艺丰",
			DepositTime:       1551004856,
			OrderNo:           "no4596609",
		},
		{
			ArtworkNo:         "no458990",
			TransactionDate:   1551004856,
			TransactionAmount: "82000万",
			Seller:            "王五",
			Buyer:             "钱**",
			AssuranceUnit:     "弘艺丰",
			DepositTime:       1551004999,
			OrderNo:           "no4596609",
		},
		{
			ArtworkNo:         "no458990",
			TransactionDate:   1551804800,
			TransactionAmount: "80000万",
			Seller:            "张三",
			Buyer:             "李*",
			AssuranceUnit:     "弘艺丰",
			DepositTime:       1551804856,
			OrderNo:           "no4590000",
		},
		{
			ArtworkNo:         "no458990",
			TransactionDate:   1551804856,
			TransactionAmount: "80000万",
			Seller:            "张三",
			Buyer:             "李*",
			AssuranceUnit:     "弘艺丰",
			DepositTime:       1551804856,
			OrderNo:           "no4590000",
		},
		{
			ArtworkNo:         "no458990",
			TransactionDate:   1551804856,
			TransactionAmount: "80000万",
			Seller:            "张三",
			Buyer:             "李*",
			AssuranceUnit:     "弘艺丰",
			DepositTime:       1551804856,
			OrderNo:           "no4590000",
		},
		{
			ArtworkNo:         "no458990",
			TransactionDate:   1551804856,
			TransactionAmount: "80000万",
			Seller:            "张三",
			Buyer:             "李*",
			AssuranceUnit:     "弘艺丰",
			DepositTime:       1551804856,
			OrderNo:           "no4590000",
		},
		{
			ArtworkNo:         "no458990",
			TransactionDate:   1551804856,
			TransactionAmount: "80000万",
			Seller:            "张三",
			Buyer:             "李*",
			AssuranceUnit:     "弘艺丰",
			DepositTime:       1551804856,
			OrderNo:           "no4590000",
		},
		{
			ArtworkNo:         "no458990",
			TransactionDate:   1551804856,
			TransactionAmount: "80000万",
			Seller:            "张三",
			Buyer:             "李*",
			AssuranceUnit:     "弘艺丰",
			DepositTime:       1551804856,
			OrderNo:           "no4590000",
		},
		{
			ArtworkNo:         "no458990",
			TransactionDate:   1,
			TransactionAmount: "80000万",
			Seller:            "张三",
			Buyer:             "李*",
			AssuranceUnit:     "弘艺丰",
			DepositTime:       1551804856,
			OrderNo:           "no4590000",
		},
	}
	fmt.Printf("%+v\n", trades)

	sort.Slice(trades, func(i, j int) bool {
		return trades[i].TransactionDate < trades[j].TransactionDate
	})

	fmt.Println("===========================================================")
	fmt.Printf("%+v\n", trades)
}
