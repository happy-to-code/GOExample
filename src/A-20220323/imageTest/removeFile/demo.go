package main

import "os"

func main() {
	err := os.Remove("E:\\20.06.16Project\\GOExample\\src\\A-20220323\\imageTest\\demo6\\dst.jpg")
	if err != nil {
		panic(err)
	}
}
