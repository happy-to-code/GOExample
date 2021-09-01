package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	var m = make(map[string]string)
	m["http://10.1.3.150:9999"] = "http://101.1.3.151:8888"

	u := "http://10.1.3.150:9999/test/jml/v1/query"
	parse, err := url.Parse(u)
	if err != nil {
		panic(err)
	}
	fmt.Println("parse:", parse)
	host := parse.Host
	host = "http://" + host
	fmt.Println("host:", host)
	replace := strings.Replace(u, host, m[host], 1)
	fmt.Println("replace:", replace)
	fmt.Printf("%T\n", replace)

	fmt.Println("--------------------")
	fmt.Println([]string{})
	fmt.Println(len([]string{}))

}
