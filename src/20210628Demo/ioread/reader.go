package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("./test.1.json")
	if err != nil {
		panic(err)
	}
	fmt.Println("file:", string(file))

}
