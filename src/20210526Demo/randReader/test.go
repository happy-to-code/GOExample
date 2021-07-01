package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		// Generate symmetric key
		key := make([]byte, 16)
		_, err := rand.Reader.Read(key)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", key)
	}
}
