package main

import "fmt"

func main() {
	var bs = []Boy{
		{Name: "xm", Age: 17},
		{Name: "xh", Age: 18},
		{Name: "xl", Age: 13},
		{Name: "xz", Age: 15},
	}
	fmt.Println(bs)
	for i, b := range bs {
		b.Id = uint64(i)
		bs[i] = b
	}
	fmt.Println(bs)

}

type Boy struct {
	Id   uint64
	Name string
	Age  int
}
