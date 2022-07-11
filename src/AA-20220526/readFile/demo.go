package main

import (
	"fmt"
	"time"
)

func main() {
	// url := "E:\\20.06.16Project\\GOExample\\src\\AA-20220526\\readFile\\img.png"
	// content, err := os.ReadFile(url)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// // [123 13 10 32 32 34 97 97 34 58 32 49 50 51 44 13 10 32 32 34 98 98 34 58 32 34 120 109 34 13 10 125]
	// fmt.Println(content)

	// 1651892056090
	// 1654652942457
	// 1654652879
	// 1654652915310432
	fmt.Println(time.Now().UnixMilli())
}
