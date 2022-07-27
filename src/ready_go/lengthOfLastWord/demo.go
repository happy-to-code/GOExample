package main

import "strings"

func lengthOfLastWord(s string) int {
	split := strings.Split(strings.TrimSpace(s), " ")
	i := len(split)
	return len(split[i-1])

}

func main() {
	var s = "Hello World"
	println(lengthOfLastWord(s))
}
