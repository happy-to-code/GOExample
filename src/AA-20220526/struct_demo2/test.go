package main

import (
	"encoding/json"
	"fmt"
)

type Girl struct {
	Name string `json:"name"`
	AGE  int    `json:"age"`
}

func main() {
	var s = `{"NAME":"John", "AGE":18}`

	var g Girl
	err := json.Unmarshal([]byte(s), &g)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", g)
	bytes, _ := json.Marshal(g)
	fmt.Println(string(bytes))

}
