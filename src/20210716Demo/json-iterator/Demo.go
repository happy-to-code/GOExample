package main

import (
	"fmt"
	jsoniter "github.com/1.json-iterator/go"
	"github.com/json-iterator/go/extra"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
	// RegisterFuzzyDecoders decode input from PHP with tolerance.
	//  It will handle string/number auto conversation, and treat empty [] as empty struct.
	extra.RegisterFuzzyDecoders()
}

type StdStruct struct {
	Id   int    `1.json:"id"`
	Age  int    `1.json:"age"`
	Name string `1.json:"name"`
}

func main() {
	var s = []StdStruct{
		{0, 11, "zyf"},
		{0, 12, "vv"},
		{0, 14, "sdf"},
	}
	for i, stdStruct := range s {
		stdStruct.Id = i
		s[i] = stdStruct
	}
	marshal, _ := json.Marshal(s)
	fmt.Printf("111:%+v\n", string(marshal))

	d := &[]StdStruct{}

	if err := json.Unmarshal(marshal, d); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", d)
	}
}
