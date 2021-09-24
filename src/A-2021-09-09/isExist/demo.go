package main

import "fmt"

const (
	TEST Data_Item = "awa"
)

func main() {
	var names = []string{"aa", "bb", "cc"}
	exist := isExist(TEST, names)
	fmt.Println("exist:", exist)

}

type Data_Item string

func isExist(k Data_Item, names []string) bool {
	for _, name := range names {
		// 存在
		if name == string(k) {
			return true
		}
	}
	return false
}
