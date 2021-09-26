package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	num := randNum()
	if num < 4 {
		for i := 0; i < 500; i++ {
			num = randNum()
			if num < 4 {
				fmt.Println("retry:", i, "--", num)
				continue
			} else {
				fmt.Println("get u:", num)
				return
			}
		}
	} else {
		fmt.Println("once")
	}
}

func randNum() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(5)
}
