package main

import "fmt"

func main() {
	var d = Demo{
		Keys: []string{""},
		Name: "xm",
	}
	fmt.Printf("%+v\n", d)
	fmt.Printf("%d\n", len(d.Keys))

}

type Demo struct {
	Keys []string
	Name string
}
