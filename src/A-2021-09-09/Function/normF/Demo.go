package main

import "fmt"

func main() {
	var b = Boy{"XiaoMing", 18}
	fmt.Println("b1:", b)
	b2 := b.testB()
	fmt.Println("b2:", b2)
	fmt.Println("=======================")
	fmt.Println("b1:", b)

}

type Boy struct {
	Name string
	Age  int
}

func (b Boy) testB() Boy {
	b.Name = "xx"
	b.Age = 17
	return b
}
