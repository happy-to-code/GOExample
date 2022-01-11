package main

import (
	"crypto/sha256"
	"fmt"
)

func GenerateSequence256(seed []byte, count int) [][]byte {
	var result [][]byte
	current := seed
	for i := 0; i < count; i++ {
		current = getHash(current)
		result = append(result, current)
	}
	return result
}

func getHash(input []byte) []byte {
	h := sha256.New()
	h.Write(input)
	return h.Sum(nil)
}

func main() {
	generateSequence256 := GenerateSequence256([]byte("HelloWorld"), 2)
	fmt.Println(generateSequence256)
	fmt.Println(string(generateSequence256[0]))

}
