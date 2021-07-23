package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcsdfds"
	suffix := strings.HasPrefix(s, "abc")
	fmt.Println(suffix)

	var AA struct {
		Name string
		Age  int
	}
	AA.Age = 18
	AA.Name = "xm"
	fmt.Println("-->", AA)

	// 	------------------

	var url = "http://172.21.43.42:60002/jml/v1/bigdata/info/query"
	split := strings.Split(url, "/jml")
	fmt.Println(">", split[0])
	replace := strings.Replace(url, split[0], "127.0.0.1:8088", 1)
	fmt.Println("replace:", replace)
	// fmt.Println(">", split[0])
	//
	// newS := "http://" + replace + split[1]
	// fmt.Println(newS)
}
