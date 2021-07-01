package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := "{\"authId\":\"1002\",\"receiver\":null,\"result\":{\"bizOrderId\":\"002\",\"status\":\"accepted\",\"comments\":\"自然人严重失信黑名单2\",\"name\":\"张三\",\"id\":\"321523189012301234\",\"creditAmount\":0,\"duration\":0,\"interest\":\"7.8\",\"product\":\"\"}}"
	m := make(map[string]interface{})
	json.Unmarshal([]byte(s), &m)

	fmt.Printf("%v\n", m)
}
