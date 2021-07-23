package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "./file/test.json"
	dstFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	s := `
		hello
		world
111,123
`
	dstFile.Write([]byte(s))
}
