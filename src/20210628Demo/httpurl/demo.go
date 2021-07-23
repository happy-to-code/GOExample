package main

import (
	"fmt"
	"net/url"
)

func main() {
	// s := "http://58.220.193.203:8088/fast-api/peasantHousehold/getDataList?beginTime=2021-06-13 00:00:00&endTime=2021-06-23 00:00:00"
	// ss := "http://58.220.193.203:8088/fast-api/peasantHousehold/getDataList?beginTime=2021-06-13%2000:00:00&endTime=2021-06-23%2000:00:00"

	escape := url.QueryEscape("2021-06-13 00:00:00")
	fmt.Println(escape)

}
