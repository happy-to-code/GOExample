package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var s = "{\"ps_1002\":\"{\\\"code\\\":\\\"1002\\\",\\\"amount\\\":200,\\\"desc\\\":\\\"加入奖励积分\\\"}\",\"ps_1001\":\"{\\\"code\\\":\\\"1001\\\",\\\"amount\\\":100,\\\"desc\\\":\\\"加入奖励积分\\\"}\"}"

	var m map[string]string

	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		panic(err)
	}

	for k, v := range m {
		fmt.Println(k, "===", v)
	}
}
