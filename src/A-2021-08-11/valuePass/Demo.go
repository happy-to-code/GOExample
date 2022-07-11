package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type Boy struct {
	Name string `1.json:"name,omitempty"`
	Age  int    `1.json:"age,omitempty"`
}

func main() {
	var b = Boy{"张三", 18}
	updateName(b)
	fmt.Println("Boy:", b)

}

func updateName(b Boy) string {
	bytes, _ := json.Marshal(b)
	fmt.Println("1：", hex.EncodeToString(bytes))
	b.Name = "ss"
	bytes1, _ := json.Marshal(b)
	fmt.Println("2: ", hex.EncodeToString(bytes1))

	return ""
}
