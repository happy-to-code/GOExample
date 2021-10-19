package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	var b = Boy{
		Name: "XM",
		Age:  18,
	}
	boyBs, _ := json.Marshal(b)

	h := sha256.New()
	h.Write(boyBs)
	fmt.Printf("%x\n", h.Sum(nil))
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().Unix())

	// 140000000000000
	// 18446744073709551615
	var d uint64 = 150000000000000
	d = d * 1000000
	fmt.Println("d:", d)
	var ss uint64 = 18446744073709551615
	// sss := ss << 6
	fmt.Println("ss:", ss)
	fmt.Println("ss:", ss/1000000)
	fmt.Println("ss:", 2<<58)

}

type Boy struct {
	Name string
	Age  int
	Sex  int
}
