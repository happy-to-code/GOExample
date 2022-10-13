package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	var s = "hello"
	v := md5V(s)
	fmt.Println(v)
	// 5d41402abc4b2a76b9719d911017c592
	// 5d41402abc4b2a76b9719d911017c592
	// 	0o	I1 +/
}

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
