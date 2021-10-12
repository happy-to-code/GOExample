package main

import "fmt"

func main() {
	var hasSendBlockHeight = 300

	var toDelList []int
	mod := hasSendBlockHeight % 100
	if mod == 0 && hasSendBlockHeight > 100 {
		for i := hasSendBlockHeight - 199; i <= hasSendBlockHeight-100; i++ {
			toDelList = append(toDelList, i)
		}
	}

	fmt.Println(toDelList)
}
