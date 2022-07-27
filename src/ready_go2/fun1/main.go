package main

import "fmt"

func main() {
	name := "xiaoming"
	f(name, sayHello)
	f(name, sayHello2)

}

func sayHello(name string) {
	fmt.Printf("%s,hello\n", name)
}
func sayHello2(name string) {
	fmt.Printf("=====>%s,hello\n", name)
}

func f(name string, ff func(s string)) {
	ff(name)
}
