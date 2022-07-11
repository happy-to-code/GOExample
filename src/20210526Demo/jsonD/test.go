package main

import (
	"encoding/json"
	"fmt"
)

type IndividualCipher struct {
	AuthID      string `1.json:"authId"`
	SubjectType string
	CipherText  string
	IdHash      string
	BankID      string
	CreateTime  string // 格式为：yyyy-MM-dd hh:mm:ss
	EndTime     string // 格式为：yyyy-MM-dd hh:mm:ss
	FileHash    string
	FilePath    string
	Holder      string
}

// Individual4C 存到合约中的个体信息
type Individual4C struct {
	Data IndividualCipher
	Sign string
}
type PersonInfo struct {
	Qlr  string `1.json:"qlr"`  // 权利人名字
	Sfzh string `1.json:"sfzh"` // 权利人身份证
}

func main() {
	var i IndividualCipher = IndividualCipher{AuthID: "ijsdf", SubjectType: "legal"}
	var ic = Individual4C{Data: i, Sign: "jjjjj"}
	fmt.Println("ic:", ic)
	fmt.Println("----------------------")
	bytes, _ := json.Marshal(ic)
	fmt.Println(string(bytes))

	var p = PersonInfo{Qlr: "individual.Subject.Name", Sfzh: "320999"}
	marshal, _ := json.Marshal(p)
	fmt.Println(string(marshal))

}
