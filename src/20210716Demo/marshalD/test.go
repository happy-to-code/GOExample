package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := `{"authId":"9a72153b5a034e2b9b1d44e98c609abc","bankId":"3Xq4b5MxauMYmvWYqJorDgtoF3uHpata7MqnmD9qnLKKRkQyVF","bankName":"同济","receiverPubkeys":["LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0NCk1Ga3dFd1lIS29aSXpqMENBUVlJS29FY3oxVUJnaTBEUWdBRUxxR0Y4SVlDdyswRmdCSXJEZWZ6N1RuVGJMWG8NClVBR3MwZG1xVzYvbGtrZjRtbWhRYmx3TVlWMzNVbU5ZeEFOcUNiR2tnUWVUc1NCUi81YTVpVWhXSGc9PQ0KLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0t","LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0NCk1Ga3dFd1lIS29aSXpqMENBUVlJS29FY3oxVUJnaTBEUWdBRWpkQ05Xd1BpM0tOUUdjdFlaaEhLTlB5SWd2RXINCjZ1Qkd5S2MrN2EwY1RYUDVBWVlpemIzUUFoWGg1WHZMVnVkTk8wUXhQbHlhM3dGd1VIREVpcG5LVGc9PQ0KLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0t"],"result":{"bizOrderId":"业务订单ID777","status":"accepted","comments":"信用良好，授信批准","name":"谢春艳","id":"321322198503142825","creditAmount":30000,"duration":36,"interest":"7.8","product":"xxx助农贷"}}`

	var uploadCreditLoan UploadCreditLoan
	err := json.Unmarshal([]byte(s), &uploadCreditLoan)
	if err != nil {
		panic(err)
	}
	fmt.Println(uploadCreditLoan)

}

type UploadCreditLoan struct {
	AuthId          string   `1.json:"authId" binding:"required"`
	BankId          string   `1.json:"bankId"`
	BanKName        string   `1.json:"bankName"`
	ReceiverPubkeys []string `1.json:"receiverPubkeys"`
	Result          Result   `1.json:"result" binding:"required,dive"`
}

type Result struct {
	BizOrderId   string `1.json:"bizOrderId"`                // 金民通业 务订单ID
	Status       string `1.json:"status" binding:"required"` // accepted - 授信审批通过 rejected - 授信审批驳回
	Comments     string `1.json:"comments"`                  // 授信审批意见
	Name         string `1.json:"name" binding:"required"`
	Id           string `1.json:"id" binding:"required"`
	CreditAmount int64  `1.json:"creditAmount"`
	Duration     int    `1.json:"duration"`
	Interest     string `1.json:"interest"`
	Product      string `1.json:"product"`
}
