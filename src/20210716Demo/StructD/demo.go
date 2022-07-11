package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var s = `
{"msg":"success","code":0,"data":null}
`
	var gridRes GridRes
	err := json.Unmarshal([]byte(s), &gridRes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", gridRes)

	// var gridRes GridRes = GridRes{
	// 	Msg: "success",
	// 	Code: 0,
	// }
	//
	// gridResBytes, _ := 1.json.Marshal(gridRes)
	// fmt.Println(string(gridResBytes))

}

type GridRes struct {
	Msg  string     `1.json:"msg"`
	Code int        `1.json:"code"`
	Data []GridInfo `1.json:"data"`
}

// GridInfo 网格数据
type GridInfo struct {
	Id                    string         `1.json:"id,omitempty"`
	HouseholdRegisterNum  string         `1.json:"householdRegisterNum,omitempty"`
	HouseholderOccupation string         `1.json:"householderOccupation,omitempty"`
	HighestEducation      string         `1.json:"highestEducation,omitempty"`
	IsMajorDiseases       string         `1.json:"isMajorDiseases,omitempty"`
	Village               string         `1.json:"village,omitempty"`
	IsSocialBehavior      string         `1.json:"isSocialBehavior,omitempty"`
	BusinessType          string         `1.json:"businessType,omitempty"`
	BusinessName          string         `1.json:"businessName,omitempty"`
	RealBusinessOperator  string         `1.json:"realBusinessOperator,omitempty"`
	OperationStatus       string         `1.json:"operationStatus,omitempty"`
	EffectiveLaborNum     string         `1.json:"effectiveLaborNum,omitempty"`
	OperationMode         string         `1.json:"operationMode,omitempty"`
	Area                  float64        `1.json:"area,omitempty"`
	Rent                  float64        `1.json:"rent,omitempty"`
	AnnualIncome          string         `1.json:"annualIncome,omitempty"`
	UpdateTime            string         `1.json:"updateTime,omitempty"`
	AreaCode              string         `1.json:"areaCode,omitempty"`
	Members               []MemberObject `1.json:"members"`
}

type MemberObject struct {
	MemberName             string `1.json:"memberName,omitempty"`
	MemberCardNo           string `1.json:"memberCardNo,omitempty"`
	MemberOccupation       string `1.json:"memberOccupation,omitempty"`
	MemberHighestEducation string `1.json:"memberHighestEducation,omitempty"`
	MemberMaritalStatus    string `1.json:"memberMaritalStatus,omitempty"`
	MemberContact          string `1.json:"memberContact,omitempty"`
	Relationship           string `1.json:"relationship,omitempty"`
	Sex                    string `1.json:"sex,omitempty"`
	UpdateTime             string `1.json:"updateTime,omitempty"`
}
