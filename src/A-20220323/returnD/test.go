package main

import "fmt"

func main() {
	fmt.Println("begin")
	test()
	fmt.Println("end")
}
func test() {
	fmt.Println("111")
	if true {
		fmt.Println("true")
		return
	}
	fmt.Println("2222")

}
