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
	// gridResBytes, _ := json.Marshal(gridRes)
	// fmt.Println(string(gridResBytes))

}

type GridRes struct {
	Msg  string     `json:"msg"`
	Code int        `json:"code"`
	Data []GridInfo `json:"data"`
}

// GridInfo 网格数据
type GridInfo struct {
	Id                    string         `json:"id,omitempty"`
	HouseholdRegisterNum  string         `json:"householdRegisterNum,omitempty"`
	HouseholderOccupation string         `json:"householderOccupation,omitempty"`
	HighestEducation      string         `json:"highestEducation,omitempty"`
	IsMajorDiseases       string         `json:"isMajorDiseases,omitempty"`
	Village               string         `json:"village,omitempty"`
	IsSocialBehavior      string         `json:"isSocialBehavior,omitempty"`
	BusinessType          string         `json:"businessType,omitempty"`
	BusinessName          string         `json:"businessName,omitempty"`
	RealBusinessOperator  string         `json:"realBusinessOperator,omitempty"`
	OperationStatus       string         `json:"operationStatus,omitempty"`
	EffectiveLaborNum     string         `json:"effectiveLaborNum,omitempty"`
	OperationMode         string         `json:"operationMode,omitempty"`
	Area                  float64        `json:"area,omitempty"`
	Rent                  float64        `json:"rent,omitempty"`
	AnnualIncome          string         `json:"annualIncome,omitempty"`
	UpdateTime            string         `json:"updateTime,omitempty"`
	AreaCode              string         `json:"areaCode,omitempty"`
	Members               []MemberObject `json:"members"`
}

type MemberObject struct {
	MemberName             string `json:"memberName,omitempty"`
	MemberCardNo           string `json:"memberCardNo,omitempty"`
	MemberOccupation       string `json:"memberOccupation,omitempty"`
	MemberHighestEducation string `json:"memberHighestEducation,omitempty"`
	MemberMaritalStatus    string `json:"memberMaritalStatus,omitempty"`
	MemberContact          string `json:"memberContact,omitempty"`
	Relationship           string `json:"relationship,omitempty"`
	Sex                    string `json:"sex,omitempty"`
	UpdateTime             string `json:"updateTime,omitempty"`
}
