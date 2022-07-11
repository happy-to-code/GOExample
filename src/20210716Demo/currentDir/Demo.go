package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
)

func main() {
	var filepath string

	sysType := runtime.GOOS
	if sysType == "windows" {
		// windows系统
		filepath = "1.1.json"
	} else {
		filepath = "./1.1.json"
	}
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
