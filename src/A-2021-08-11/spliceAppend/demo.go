package main

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/1.json-iterator/go"
)

func main() {
	var events1 = []Event{
		{
			Topic:    "txevent",
			SubTopic: "onchain",
			LedgerID: "sdfddsf",
		},
	}

	var events2 = []Event{
		{
			Topic:    "txevent2",
			SubTopic: "onchain2",
			LedgerID: "sdfddsf",
		},
		{
			Topic:    "scevent3",
			SubTopic: "onchain3",
			LedgerID: "sdfddsf",
		},
	}

	newEvent := append(events2, events1...)

	fmt.Println("-->", newEvent)
	for i, v := range newEvent {
		v.ID = uint64(i)
		newEvent[i] = v
	}
	fmt.Println("-->", newEvent)

	marshal, _ := json.Marshal(&newEvent)

	var evs []Event

	jsoniter.Unmarshal(marshal, &evs)
	fmt.Println("-->", evs)

}

type Event struct {
	ID          uint64      `1.json:"id"` // 事件的ID
	Topic       string      `1.json:"topic"`
	SubTopic    string      `1.json:"subTopic"`
	Date        string      `1.json:"date"`        // 事件发生时间
	LedgerID    string      `1.json:"ledgerID"`    // 触发事件的链
	BlockHeight uint64      `1.json:"blockHeight"` // 触发事件的区块
	Port        int         `1.json:"port"`        // 发送事件节点的端口
	Client      string      `1.json:"client"`      // 事件发起方——sdk ID
	EventMsg    interface{} `1.json:"eventMsg"`    // 事件消息
}
