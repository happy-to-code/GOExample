package main

import "fmt"

var hasPushedBlockHeight uint64

func main() {
	fmt.Println(hasPushedBlockHeight)
	hasPushedBlockHeight = 50

	var a = hasPushedBlockHeight

	hasPushedBlockHeight = 60
	fmt.Println("a:", a)
	fmt.Println("hasPushedBlockHeight:", hasPushedBlockHeight)

}
