package main

import (
	"encoding/json"
	"fmt"
)

type MaPDemo struct {
	M map[string]string
	N string
}

func main() {
	m := make(map[string]string)
	m["a"] = "aa"
	m["b"] = "bb"
	fmt.Println("m:", m)

	var mapDemo = MaPDemo{
		M: m,
		N: "xm",
	}
	m1 := mapDemo.M
	marshal, _ := json.Marshal(m1)

	var m3 map[string]string
	json.Unmarshal(marshal, &m3)

	fmt.Println("m2  1", string(marshal))

	changeMap(m3)
	fmt.Println("m3  2", m3)

	fmt.Println("_____________________")
	fmt.Println(m1)

}

func changeMap(m map[string]string) {

	m["a"] = "a1a"
}

func DeepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = DeepCopy(v)
		}

		return newSlice
	}

	return value
}
