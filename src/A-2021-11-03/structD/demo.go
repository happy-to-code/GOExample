package main

import (
	"encoding/json"
	"fmt"
)

// IndividualCipher 存到合约中的个体信息
type IndividualCipher struct {
	AuthID      string      `1.json:"authId"`
	SubjectType string      `1.json:"subjectType"`
	CipherText  string      `1.json:"cipherText"`
	CipherKeys  []CipherKey `1.json:"cipherKeys"`
	IdHash      string      `1.json:"idHash"`
	BankID      string      `1.json:"bankId"`
	CreateTime  string      `1.json:"createTime"` // 格式为：yyyy-MM-dd hh:mm:ss
	EndTime     string      `1.json:"endTime"`    // 格式为：yyyy-MM-dd hh:mm:ss
	FileHash    string      `1.json:"fileHash"`
	// FilePath    string      `1.json:"filePath"`
	Holder string `1.json:"holder"`
}
type CipherKey struct {
	Id        string `1.json:"id"`
	Keycipher string `1.json:"keycipher"` // 密钥
}

// Individual4C 存到合约中的个体信息
type Individual4C struct {
	Data IndividualCipher `1.json:"data"`
	Sign string           `1.json:"Sign"`
}

func main() {
	s := `{
"authId": "a55e4abfda8d4030bc52f400aaf30f50",
"subjectType": "natural",
"bankId": "4Df1RTG114T2y6y2jtXW5BTCGbRPXu4giXJnuTm86BQkk1NFKU",
"BankName": "高邮兴福村镇银行",
"endTime": "9999-12-31 00:00:00",
"fileHash": "398895E76C339B918C4F52E4A219EDA65C1B2A11E1993CA572951CB925BA3437",
"subject": {
"id": "",
"name": "",
"code": ""
}
}`
	authorizeBytes := []byte(s)
	var individual4C Individual4C
	json.Unmarshal(authorizeBytes, &individual4C)

	fmt.Println(individual4C)
}
