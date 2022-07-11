package main

import "fmt"

var bb *Boy

func main() {
	var b = Boy{"xm"}
	bb = &b

	fmt.Println("bb:", bb)
	fmt.Println("&bb:", &bb)
	bb.Name = "xxxxx"
	fmt.Println("-->&bb:", bb)

	dbTmp := *bb
	dbTmp2 := &dbTmp

	fmt.Println("dbTmp2:", dbTmp2)
	fmt.Println("&dbTmp2:", &dbTmp2)
}

type Boy struct {
	Name string
}
