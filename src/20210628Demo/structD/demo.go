package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var qr = QueryRecord{
		CardId:     "32098715485664",
		ChainHash:  "txid",
		Id:         "requestId",
		InquirerId: "bankId",
		DataSourceAndItemId: []SourceAndItem{
			{DataSourceId: "111", DataItemId: []string{"112", "113"}},
			{DataSourceId: "222", DataItemId: []string{"222", "223"}},
		},
	}

	marshal, _ := json.Marshal(qr)
	fmt.Println(string(marshal))
}

type QueryRecord struct {
	CardId              string          `json:"cardId"`     // 被查询人身份标识
	ChainHash           string          `json:"chainHash"`  // 链上哈希
	Id                  string          `json:"id"`         // 请求的ID，requestId
	InquirerId          string          `json:"inquirerId"` // 查询方机构id
	DataSourceAndItemId []SourceAndItem `json:"dataSourceAndItemId"`
}

// type DataSourceAndItemIds struct {
// 	DataSourceAndItemId []SourceAndItem `json:"dataSourceAndItemId"`
// }

type SourceAndItem struct {
	DataSourceId string   `json:"dataSourceId"` // 数据源id
	DataItemId   []string `json:"dataItemId"`   // 数据项id
}
