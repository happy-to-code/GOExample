package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

func genHash(data []byte) string {
	h := sha256.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// 9f26a8908e54a714557454d8a8cf3188dae364af3b49535be83561ff5b1e106e
// 9f26a8908e54a714557454d8a8cf3188dae364af3b49535be83561ff5b1e106e
// 3574e99f4ea98e28e5d631166e007ab4b56a1bc0272ad2757434f246915f2f22
// 3574e99f4ea98e28e5d631166e007ab4b56a1bc0272ad2757434f246915f2f22
func main() {
	var s = "123456789abcde"
	hash := genHash([]byte(s))
	fmt.Println("hash:", hash)

	// 第一种调用方法
	sum := sha256.Sum256([]byte(s))
	fmt.Printf("%x\n", sum)
	fmt.Println("111:", hex.EncodeToString(sum[:]))

	// 第二种调用方法
	h := sha256.New()
	h.Write([]byte(s))
	fmt.Printf("%x\n", h.Sum(nil))
	fmt.Println("222:", hex.EncodeToString(h.Sum(nil)))

	fmt.Println(time.Now().UnixMilli())
	fmt.Println(strconv.Itoa(int(time.Now().UnixMilli())))
	// 1655715867604
}
