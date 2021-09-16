package main

import (
	"fmt"
	"strconv"
)

type Boy struct {
	Name string
	Age  int
}

func main() {
	var boy = Boy{"XM", 11}
	fmt.Println("boy:", boy)

	var boyBack = boy
	boy.Name = "HHHH"
	fmt.Println("boy2:", boy)

	fmt.Println("boyBack:", boyBack)

	fileName := boy.Name + "/" + strconv.Itoa(int(boy.Age)) + "_bk"
	fmt.Println(fileName)
}

func test(b Boy) Boy {
	b.Name = ""
	return b
}
