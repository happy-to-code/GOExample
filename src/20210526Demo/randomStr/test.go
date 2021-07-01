package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(createRandomString(10))
	// }

	var iv = make([]byte, 16)
	fmt.Println(iv)
}
func createRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyz1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}
