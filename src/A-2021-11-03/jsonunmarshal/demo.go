package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 	var errS = `[{"id":0,"topic":"scevent","subTopic":"addAuthorize","date":"2021-12-23 14:37:08","ledgerID":"hgxn0wvvki","blockHeight":1929731,"port":0,"txList":[{"txHash":"g6gW
	// JVdEpVF79I77eGPerYLFI3JZNv0zn0T/pb9XgJU="}],"client":"","eventMsg":"{\"data\":{\"authId\":\"f5ba1d0716144b6c98cd623446bff55a\",\"subjectType\":\"natural\",\"cipherText\
	// ":\"bfc511610905c61f058a4b6dcd706085cc8a69512d08eee2eed4bacd7f6ec4630a64e37c0f37baae3dc7b3c04f3c275ae172ed99fa34e2f990b0e535398a2ea1\",\"cipherKeys\":[{\"id\":\"47Pfpyt
	// B\",\"keycipher\":\"047581bda370cdaf1d15f0a4c3950f53d8035a23d8d6ca4484d66fd4a3f5a144212c64180d1d74f547cd38b31e94e622f2f0a57eff68a525daaed2e45c92266d4868f9211f45ff9fdd07
	// eda1d91527ff3ae0d5729dd4e2f5a31ad855ab313a82e4f9dd15c86a2688aa879a1a40ef1cc37b\"},{\"id\":\"39FGv5VA\",\"keycipher\":\"04d538d59bc06aba0467bda8c7684def216da17e477c3b0a3
	// 51053aab040800b348cd9ef92acdd7ea814a153a480a05e839c47fd372583c30b3920358e7851bef7d5ec14c7bfba8bb2029dff730f7e5c2e40c8f210d725cdc251ba1f9939562e92326340580ef81e7f4adea76
	// e1732a683\"}],\"idHash\":\"cfc284fd118ae8f1f1f23852e3c298fac143da8a20fb03c5da5b607cf3d25aeb\",\"bankId\":\"39FGv5VA67hgKqBLQ8B8g8Aay8su9Y1rCrNJSm3x6ki9BjzDRv\",\"create
	// Time\":\"2021-12-23 14:38:39\",\"endTime\":\"2022-01-22 14:38:39\",\"fileHash\":\"861725dbfeed8a7b0c6d6ef6125e31204be92ad5c5787bca2ec74959182987bc\",\"holder\":\"39FGv5
	// VA67hgKqBLQ8B8g8Aay8su9Y1rCrNJSm3x6ki9BjzDRv\"},\"Sign\":\"3046022100edc86d8d21e10582594585d372e57df44175b1403ee5278b7a3a16f7b46bb4620221008b4a9c6c8559531a3cd8326d79e97
	// 18b66cef5bbcc4e0bc3308d52db45a2a6f5\"}FGF83a816255744a5517bf48efb7863dead82c523725936fd339f44ffa5bf578095FGF1640241519640"},{"id":1,"topic":"txevent","subTopic":"onchai
	// n","date":"2021-12-23 14:38:39","ledgerID":"hgxn0wvvki","blockHeight":1929731,"port":0,"txList":[{"txHash":"g6gWJVdEpVF79I77eGPerYLFI3JZNv0zn0T/pb9XgJU="}],"client":"bd
	// 99a42a98d8f15180e02fb83cd94c9426819d3bee006bca0bca29d7cb11f0df45a87cf883cfd93f9371fa5b510f5191586922371b9fe8d307e4d78ddf6b0521","eventMsg":"SUCCESS"}]`

	var rightS = `
[{"id":0,"topic":"scevent","subTopic":"addAuthorize","date":"2021-12-23 14:37:04","ledgerID":"hgxn0wvvki","blockHeight":1929729,"port":0,"txList":[{"txHash":"og/wQs8U3Z
SAfw7HDVHaU1d4EKDf3gQwewkLiL9noTk="}],"client":"","eventMsg":"{\"data\":{\"authId\":\"6c93ccc326e04674aff819e726dee524\",\"subjectType\":\"natural\",\"cipherText\":\"2e
62a1b98208fc1d8c4d8cd0d08cf21bcef69c8ef3ca1a457a47233f9d3879e4db8016a2bc7a430b25061b4f7892b3cb4713aac86d6a7f9410308674b652dff6\",\"cipherKeys\":[{\"id\":\"47PfpytB\",\"
keycipher\":\"040e0f022deee78de134dabd085a582d4981099b5ad0bf8ae8fec4d9f81efa6a91d42f3a70f130c78ed65309692774da577861b9c1119f24c2f5d8d330d50d970a09a45555259e7ed606ac3769
6d39ddc3021116e1bd83518e009c8ec0e2cb8c780afcf0280e63f9d8e446118700bd5da4\"},{\"id\":\"39FGv5VA\",\"keycipher\":\"04a6a3a4a5bf17ace054628da15cde5e35214859c3df5a27786aece
8ca96654a6d2ad35a44d0f15276157effaeecd87395d003ce9a9a4fe9a48223061022ae0514b3eb3df3222cc6ccda818e55c3b46a255d181c81955c3357494ddf21554540b49052fda62e8ad4b1ea10a2a30ef85
0a5\"}],\"idHash\":\"073e270f2145a8061526cb99683f322d68a4c6896aa716281242938f0d535b48\",\"bankId\":\"39FGv5VA67hgKqBLQ8B8g8Aay8su9Y1rCrNJSm3x6ki9BjzDRv\",\"createTime\"
:\"2021-12-23 14:38:35\",\"endTime\":\"2022-01-22 14:38:35\",\"fileHash\":\"4644f35e9af79946c21d7d45768e2c87740935710ba18172c8d5113decd53bac\",\"holder\":\"39FGv5VA67hg
KqBLQ8B8g8Aay8su9Y1rCrNJSm3x6ki9BjzDRv\"},\"Sign\":\"3045022100e3560a4a6af2fbbbeafbfb4a5f2fc3670e39e0c8f2bf226b503a15b2c65fef79022002d3ced27c9d68c4df8e34d6182128589a0d2
5044bcc7e946f6d55b47aa7c990\"}FGFa20ff042cf14dd94807f0ec70d51da53577810a0dfde04307b090b88bf67a139FGF1640241515187"},{"id":1,"topic":"txevent","subTopic":"onchain","date
":"2021-12-23 14:38:35","ledgerID":"hgxn0wvvki","blockHeight":1929729,"port":0,"txList":[{"txHash":"og/wQs8U3ZSAfw7HDVHaU1d4EKDf3gQwewkLiL9noTk="}],"client":"bd99a42a98
d8f15180e02fb83cd94c9426819d3bee006bca0bca29d7cb11f0df45a87cf883cfd93f9371fa5b510f5191586922371b9fe8d307e4d78ddf6b0521","eventMsg":"SUCCESS"}]
`

	var events []Event
	err := json.Unmarshal([]byte(rightS), &events)
	if err != nil {
		panic(err)
	}
	fmt.Printf("events:%+v\n", events)
}

type Type string
type SubType string
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
}
type TxInfo struct {
	TxHash []byte `1.json:"txHash,omitempty"` // 触发事件的交易
	// TxIndex uint   `1.json:"txIndex,omitempty"` // 触发事件的交易索引
}
