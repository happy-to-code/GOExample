package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var s = `{"state":200,"code":"success","message":"successdd","data":"fa641fc661bc7decab978c07a71e4ed635e10d427d5efbc464db7ee86ead190b"}`
	var req StoreNotaryChainRes
	err := json.Unmarshal([]byte(s), &req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", req)

}

type StoreNotaryChainRes struct {
	State   int    `json:"state"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
