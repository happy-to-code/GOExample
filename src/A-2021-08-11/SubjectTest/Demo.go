package main

import (
	"encoding/json"
	"fmt"
)

type Boy struct {
	Name string
	Age  int
}

func main() {
	var boy = Boy{"XM", 11}
	s := test(boy)
	fmt.Println("s:", s)
	fmt.Println("boy:", boy)
	boyB, _ := json.Marshal(boy)
	fmt.Println("boyB:", string(boyB))
}

func test(b Boy) string {
	b.Name = ""
	marshal, _ := json.Marshal(b)
	return string(marshal)
}
