package main

import "fmt"

type Animal interface {
	Eat()
	Yell()
}

type Dog struct {
}

func (d Dog) Eat() {
	fmt.Println("吃骨头")
}

func (d Dog) Yell() {
	fmt.Println("汪汪汪")
}

type Cat struct {
}

func (c Cat) Eat() {
	fmt.Println("吃鱼鱼")
}

func (c Cat) Yell() {
	fmt.Println("喵喵喵")
}

func main() {
	var s = "Hello World"
	fmt.Println(s)
	var animal Animal = new(Dog)
	animal.Eat()
	animal.Yell()

	fmt.Println("---------------------")
	animal = new(Cat)
	animal.Eat()
	animal.Yell()
}
