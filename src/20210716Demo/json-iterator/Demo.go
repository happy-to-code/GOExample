package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
	// RegisterFuzzyDecoders decode input from PHP with tolerance.
	//  It will handle string/number auto conversation, and treat empty [] as empty struct.
	extra.RegisterFuzzyDecoders()
}

type StdStruct struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func main() {
	var s = StdStruct{11, "zyf"}
	marshal, _ := json.Marshal(s)

	d := &StdStruct{}

	if err := json.Unmarshal(marshal, d); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", d)
		fmt.Printf("%T\n", d.Age)
	}
}
