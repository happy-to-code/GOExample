package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

const json1 = `{
  "names": 
    {
      "name": "十五",
      "age" : 1
    }
}`

func main() {
	value1 := gjson.Get(json1, "names.name")
	fmt.Println(value1.String())

	value2 := gjson.Get(json1, "names.age")
	fmt.Println(value2.Int())
}
