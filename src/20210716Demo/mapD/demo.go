package main

import "fmt"

func main() {

	m := make(map[string]interface{})
	m["a"] = "AA"
	m["b"] = 12
	m["c"] = false
	for k, v := range m {
		fmt.Println(k, "--", v)
	}

	delete(m, "a")

	fmt.Println("---------------------")
	for k, v := range m {
		fmt.Println(k, "--", v)
	}

	bytes, _ := test()
	fmt.Println(len(bytes))
}

func test() ([]byte, error) {
	return nil, nil
}
