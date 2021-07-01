package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	s := "hahhhyes"
	toString := hex.EncodeToString([]byte(s))
	fmt.Println(toString)

	fmt.Println("==================")
	decodeString, _ := hex.DecodeString(toString)
	fmt.Println(string(decodeString))
}
