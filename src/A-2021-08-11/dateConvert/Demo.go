package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// 2020年9月18日  --> 2019-06-26
	// var date1 = "2020年9月18日"
	// replace := strings.Replace(date1, "年", "-", -1)
	// replace1 := strings.Replace(replace, "月", "-", -1)
	// replace2 := strings.Replace(replace1, "日", "", -1)
	// fmt.Println(replace2)

	// timestamp, err := time.Parse("2006-01-02-", date1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(timestamp.Unix())
	// fmt.Println(timestamp.Month())
	// fmt.Println(timestamp.Format("2006-01-02"))

	// loc, _ := time.LoadLocation("Local")
	// t1, _ := time.ParseInLocation("2006-01-02 ", "2020年9月18日", loc)
	// fmt.Println(t1.Unix())

	// yyyyMM
	var s = " 2020年10 月"

	s = strings.TrimSpace(s)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "年", "", -1)
	s = strings.Replace(s, "月", "", -1)
	fmt.Println("s:", s)
	if len(s) == 5 {
		sFront := s[:4]
		sEnd := s[4:]
		fmt.Println(sFront, "-", sEnd)
		s = sFront + "0" + sEnd
	}
	fmt.Println("sss:", s)

	timeStr := time.Unix(time.Now().Unix(), 0).Format("2006-01-02--15-04-05")
	fmt.Println(timeStr)

}
