package main

import "fmt"

// "subject" - 01主体
// "account" - 02账户
// "product" - 03产品
// "transactionreport" - 04交易报告
// "registration" - 05登记
// "settlement" - 06资金结算
// "infodisclosure" - 07信批
func main() {
	var typeList = []string{"subject",
		"account",
		"product",
		"transactionreport",
		"registration",
		"settlement",
		"infodisclosure"}

	var temp1 = Temp{"11", "aa", "subject"}
	// var temp11 = Temp{"111", "aaa", "subject"}
	var temp2 = Temp{"股东账号", "bb", "account"}
	var temp22 = Temp{"资金账号", "bbb", "account"}
	// var temp3 = Temp{"33", "cc", "product"}
	// var temp4 = Temp{"44", "dd", "transactionreport"}
	// var temp44 = Temp{"444", "ddd", "transactionreport"}
	// var temp5 = Temp{"55", "ee", "registration"}
	// var temp6 = Temp{"66", "ff", "settlement"}
	// var temp7 = Temp{"77", "gg", "infodisclosure"}

	var list []Temp
	// list = append(list, temp3)
	list = append(list, temp2)
	list = append(list, temp22)
	// list = append(list, temp5)
	// list = append(list, temp4)
	// list = append(list, temp7)
	// list = append(list, temp6)
	// list = append(list, temp11)
	list = append(list, temp1)
	// list = append(list, temp44)
	fmt.Println(list)

	var newList []Temp
	for _, ty := range typeList {
		tyList := findAll(ty, list)
		if len(tyList) > 0 {
			newList = append(newList, tyList...)
		}
	}
	fmt.Println(newList)
}

func findAll(ty string, list []Temp) (temps []Temp) {
	for _, temp := range list {
		if temp.Type == ty {
			temps = append(temps, temp)
		}
	}
	return
}

type Temp struct {
	Name string
	File string
	Type string
}
