package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var s = "hello,word"
	fmt.Printf("%x\n", s)
	encodeToString := hex.EncodeToString([]byte(s))
	fmt.Println(encodeToString)

	rand.Seed(time.Now().Unix())
	s2 := []string{"one", "two", "three"}[rand.Intn(3)]
	fmt.Println(s2)
}
