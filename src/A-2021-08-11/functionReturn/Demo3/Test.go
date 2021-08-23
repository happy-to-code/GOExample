package main

import "fmt"

func hello() {
	fmt.Println("hello")
}

func say(f func()) {
	f()
}

func main() {
	f := hello
	say(f)
}
