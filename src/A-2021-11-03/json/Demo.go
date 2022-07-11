package main

import (
	"encoding/json"
	"fmt"
)

// Type 事件类型
type Type string

// SubType 事件子类型
type SubType string

type TxInfo struct {
	TxHash  []byte `1.json:"txHash,omitempty"`  // 触发事件的交易
	TxIndex uint   `1.json:"txIndex,omitempty"` // 触发事件的交易索引

}
type Event struct {
	ID          uint64      `1.json:"id"` // 事件的ID
	Topic       Type        `1.json:"topic"`
	SubTopic    SubType     `1.json:"subTopic"`
	Date        string      `1.json:"date"`        // 事件发生时间
	LedgerID    string      `1.json:"ledgerID"`    // 触发事件的链
	BlockHeight uint64      `1.json:"blockHeight"` // 触发事件的区块
	Port        int         `1.json:"port"`        // 发送事件节点的端口
	TxList      []TxInfo    `1.json:"txList"`      // 交易信息
	Client      string      `1.json:"client"`      // 事件发起方——sdk ID
	EventMsg    interface{} `1.json:"eventMsg"`    // 事件消息
	// Dump     bool      //持久化
}

func main() {
	var s = `{"id":0,"topic":"scevent","subTopic":"addAuthorize","date":"2021-11-17 16:27:40","ledgerID":"fkqqshd6r9","blockHeight":5841,"port":0,"txList":[{"txHash":"z2k9uq+
61sDrlK0mWwstM/PLa1hbvk6iR/fXt+zD8DI="}],"client":"","eventMsg":"{\"data\":{\"authId\":\"0c831c6766ce4d1db0d894e853fc97cc\",\"subjectType\":\"natural\",\"cipherText\":\
"50ceb6cc7b5f95adeb41a9bc75232e1c99e30824bb3b6280dec30cbde2bb8c61313bcbfe3da16610dabdc725ce9051d5d320265b16f762b1ac1769889354ffa7\",\"cipherKeys\":[{\"id\":\"3Fipu86B\"
,\"keycipher\":\"046fdcae37a3a44a297e5cd461fc0197e24886b6cfe85907b1f254a8f34436797f7937e4f27e0828fd44244ea014015d841b7240e71a5db2669253a7bc7cffd8d68a16bbe2dc993e55fea
80cfab4d00be5e81e3f71f2baaed5a993ef54fd8df2082981ca283fb7de222508aa76006cb48c\"},{\"id\":\"4Df1RTG1\",\"keycipher\":\"04b6318e51dd8047b2fc5eafbf203fbac0abd916cfc5e66a
cfcf7f9500333e6205bbb35738514e85776f5e5a95211eaee1b048f468b81f8e44d60fea05055d2aaf1e1259a19ded62be12c7237f17c96c9fd781d5aad69662c0d51b36c9103c9c6428aa81b0c1202abf9e4b
,\"keycipher\":\"046fdcae37a3a44a297e5cd461fc0197e24886b6cfe85907b1f254a8f34436797f7937e4f27e0828fd44244ea014015d841b7240e71a5db2669253a7bc7cffd8d68a16bbe2dc993e55fea80
cfab4d00be5e81e3f71f2baaed5a993ef54fd8df2082981ca283fb7de222508aa76006cb48c\"},{\"id\":\"4Df1RTG1\",\"keycipher\":\"04b6318e51dd8047b2fc5eafbf203fbac0abd916cfc5e66acfcf
7f9500333e6205bbb35738514e85776f5e5a95211eaee1b048f468b81f8e44d60fea05055d2aaf1e1259a19ded62be12c7237f17c96c9fd781d5aad69662c0d51b36c9103c9c6428aa81b0c1202abf9e4b8d78c5
b1bd6c\"}],\"idHash\":\"865b8da1711e0b2ebb9358b53c4b30dd08331ead22bae77753364d0cbc0c4072\",\"bankId\":\"4Df1RTG114T2y6y2jtXW5BTCGbRPXu4giXJnuTm86BQkk1NFKU\",\"createTim
e\":\"2021-11-17 16:28:29\",\"endTime\":\"2021-11-30 23:59:59\",\"fileHash\":\"sdds1c7a6305fa3b2a979bf81d760aec3fca867\",\"holder\":\"4Df1RTG114T2y6y2jtXW5BTCGbRPXu4giX
JnuTm86BQkk1NFKU\"},\"Sign\":\"3046022100c52b55a6fc0455f37979697ee69577ae68794ea726a3f63432c51c0542110355022100fbe480d90d8bf206b13f556a6d4efe77f87961a9ede0e47b341b5cd01
ee287fa\"}FGFcf693dbaafbad6c0eb94ad265b0b2d33f3cb6b585bbe4ea247f7d7b7ecc3f032FGF1637137709772"}`
	eventMsg := Event{}
	err := json.Unmarshal([]byte(s), &eventMsg)
	if err != nil {
		panic(err)
	}
	fmt.Println(eventMsg)

}
