package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	// 692efc8f49ac4291c909f71470b5cb85
	s := "123456"
	hash := md5.New()
	hash.Write([]byte("sixrsixr"))
	hash.Write([]byte(s))
	hash.Write([]byte(s[:2]))
	hash.Write([]byte(s[4:]))
	hash.Write([]byte("SIXRSIXR"))
	toString := hex.EncodeToString(hash.Sum(nil))
	fmt.Println(toString)

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(currentTime)
}

// 0xcedc39708cca0088ea1cd3cbed346f52bb878d80
