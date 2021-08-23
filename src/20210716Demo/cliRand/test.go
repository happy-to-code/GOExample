package main

import (
	crand "crypto/rand"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	var seed int64
	var sb [4]byte
	crand.Read(sb[:])
	seed = int64(time.Now().Nanosecond())<<32 |
		int64(sb[0])<<24 | int64(sb[1])<<16 |
		int64(sb[2])<<8 | int64(sb[3])
	cliRand = rand.New(rand.NewSource(seed))
}

var cliRand *rand.Rand

func main() {
	sprint := fmt.Sprint(cliRand.Int63())

	fmt.Println(sprint)

}
