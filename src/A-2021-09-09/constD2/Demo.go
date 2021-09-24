package main

import "fmt"

type Data_Item string

const (
	ASSET_NA_HOUSE   Data_Item = "房屋信息"
	ASSET_NA_VEHICLE Data_Item = "车辆信息"
)

func main() {
	fmt.Println(ASSET_NA_HOUSE)
	fmt.Println(ASSET_NA_VEHICLE)

}
