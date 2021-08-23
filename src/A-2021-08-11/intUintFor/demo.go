package main

import "fmt"

func main() {
	var hUint uint64 = 13
	var hasSend int = 11

	if int(hUint) > hasSend {
		fmt.Println("-----------")
		for i := hasSend + 1; i <= int(hUint); i++ {
			fmt.Println(i)
		}
	}
}
