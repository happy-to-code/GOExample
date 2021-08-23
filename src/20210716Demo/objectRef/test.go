package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type Boy struct {
	Name string
	Age  int
}

func main() {
	var b = Boy{"XM", 18}

	fmt.Println("toString1:", b)

	var bcopy = b
	chName(bcopy)
	fmt.Println("bcopy:", bcopy)

	fmt.Println("toString2:", b)
}

func chName(bcopy Boy) {
	bcopy.Name = "xj"
	fmt.Println(bcopy)
	marshal, _ := json.Marshal(bcopy)
	toString := hex.EncodeToString(marshal)
	fmt.Println("toString:", toString)
}
