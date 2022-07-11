package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func main() {
	var s = "helloworld"
	fmt.Println("1-->", sha128([]byte(s))) // 6adfb183a4a2c94a2f92dab5ade762a47889a5a1
	fmt.Println("2-->", sha_256([]byte(s)))
	fmt.Println("3-->", sha_512([]byte(s)))

}

func sha128(by []byte) string {
	w := sha1.New()
	w.Write(by)
	result := w.Sum(nil)
	return hex.EncodeToString(result)
}

func sha_256(by []byte) string {
	w := sha256.New()
	w.Write(by)
	result := w.Sum(nil)
	return hex.EncodeToString(result)
}
func sha_512(by []byte) string {
	w := sha512.New()
	w.Write(by)
	result := w.Sum(nil)
	return hex.EncodeToString(result)
}
