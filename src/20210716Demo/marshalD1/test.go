package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var u = Upload{"AA123", "bankId345", "{\"key\":\"123\"}"}

	var uAndHash = UploadAndHash{u, "hash"}

	fmt.Println(uAndHash)

	marshal, _ := json.Marshal(uAndHash)
	fmt.Println("--->", string(marshal))

}

type Upload struct {
	AuthId   string      `1.json:"auth_id,omitempty"`
	BankId   string      `1.json:"bank_id,omitempty"`
	EventMsg interface{} `1.json:"event_msg,omitempty"`
}

type UploadAndHash struct {
	Upload `1.json:"upload"`
	Hash   string `1.json:"hash"`
}
