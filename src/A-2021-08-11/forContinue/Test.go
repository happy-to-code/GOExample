package main

import "fmt"

func main() {
	var s1 = []string{"AA", "BB", "CC"}
	var s2 = []string{"aaa", "bbb", "ccc"}

	for _, out := range s1 {
		if out == "BB" {
			continue
		}
		fmt.Println("=============>", out)
		for _, in := range s2 {
			if in == "bbb" {
				continue
			}
			fmt.Println("----------->", in)
		}
	}
}
