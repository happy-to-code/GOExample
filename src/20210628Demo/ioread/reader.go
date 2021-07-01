package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("./test.json")
	if err != nil {
		panic(err)
	}
	fmt.Println("file:", string(file))

}
