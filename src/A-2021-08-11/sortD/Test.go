package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

var s = `[
{
"JNRQ": "2020-06-20 00:00:00",
"BYSJJEHJ": 0,
"JFNY": "202006",
"TGDWQC": "扬州市人力资源与社会保障局",
"BYYJJEHJ": 0,
"TYSHXYDM": "9001993",
"JBJGQC": "扬州人社局",
"ZZJGDM": "913210007514485514",
"JGQC": "扬州兰庭物业服务有限公司",
"DWJFJS": 404160,
"XXYTBMQC": "扬州社保局",
"TGRQ": "2020-07-01 05:04:53",
"LB": "410",
"CBRS": 120
},
{
"JNRQ": "2020-06-18 00:00:00",
"BYSJJEHJ": 2020.8,
"JFNY": "202006",
"TGDWQC": "扬州市人力资源与社会保障局",
"BYYJJEHJ": 2020.8,
"TYSHXYDM": "9001993",
"JBJGQC": "扬州人社局",
"ZZJGDM": "913210007514485514",
"JGQC": "扬州兰庭物业服务有限公司",
"DWJFJS": 404160,
"XXYTBMQC": "扬州社保局",
"TGRQ": "2020-07-01 05:04:53",
"LB": "210",
"CBRS": 120
},
{
"JNRQ": "2020-06-27 00:00:00",
"BYSJJEHJ": 32332.8,
"JFNY": "202006",
"TGDWQC": "扬州市人力资源与社会保障局",
"BYYJJEHJ": 32332.8,
"TYSHXYDM": "9001993",
"JBJGQC": "扬州人社局",
"ZZJGDM": "913210007514485514",
"JGQC": "扬州兰庭物业服务有限公司",
"DWJFJS": 404160,
"XXYTBMQC": "扬州社保局",
"TGRQ": "2020-07-01 05:04:53",
"LB": "110",
"CBRS": 120
}]`

func main() {
	var ls []LegalSSPayment
	err := json.Unmarshal([]byte(s), &ls)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", ls)
	sortByDate := SortByDate(ls)
	fmt.Printf("%+v\n", sortByDate)
	fmt.Printf("%+v\n", sortByDate[:2])

}
func SortByDate(ls []LegalSSPayment) []LegalSSPayment {
	sort.Slice(ls, func(i, j int) bool {
		return ls[i].JFNY > ls[j].JFNY
	})

	return ls
}

type LegalSSPayment struct {
	JNRQ     string  `1.json:"JNRQ"`
	BYSJJEHJ float64 `1.json:"BYSJJEHJ"`
	JFNY     string  `1.json:"JFNY"`
	TGDWQC   string  `1.json:"TGDWQC"`
	BYYJJEHJ float64 `1.json:"BYYJJEHJ"`
	TYSHXYDM string  `1.json:"TYSHXYDM"`
	JBJGQC   string  `1.json:"JBJGQC"`
	ZZJGDM   string  `1.json:"ZZJGDM"`
	JGQC     string  `1.json:"JGQC"`
	DWJFJS   int     `1.json:"DWJFJS"`
	XXYTBMQC string  `1.json:"XXYTBMQC"`
	TGRQ     string  `1.json:"TGRQ"`
	LB       string  `1.json:"LB"`
	CBRS     int     `1.json:"CBRS"`
}

// ZZJGDM   string  `1.json:"company_name"` // 组织机构代码
// CBRS     int     `1.json:"persons"`      // 参保人数
// DWJFJS   float64 `1.json:"base"`         // 单位缴费基数
// JFNY     string  `1.json:"month"`        // 缴费年月
// BYYJJEHJ float64 `1.json:"payable"`      // 本月应缴金额合计（元）
// BYSJJEHJ float64 `1.json:"paid"`         // 本月实缴金额合计（元）
// LB       string  `1.json:"type"`         // 类别
