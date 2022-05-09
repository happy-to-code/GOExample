package main

import (
	"encoding/json"
	"fmt"
)

type temp struct {
	Data string `json:"data"`
}

func main() {
	var s = `{"Data":"eyJhcnRpc3RObyI6Im5vMTIzNDU2NzgiLCJhcnRpc3ROYW1lIjoi5byg5paH5Y2OIiwibmFtZVVzZWRCZWZvcmUiOiLlvKDlsI/ljY4iLCJwc2V1ZG9ueW0iOiLmn5IiLCJkb2N1bWVudFR5cGUiOiLouqvku73or4EiLCJpZE51bWJlciI6IjIzNjQ4OTU2NTYxMjM1NDc4OSIsIm5hdGl2ZVBsYWNlIjoi5rKz5Y2X5ZWG5LiYIiwiZGF0ZU9mQmlydGgiOiIxOTUyLTEyLTA0IiwiZm9ybXNPZkFydCI6IuWbveeUuyIsImdlbmRlciI6IuS/neWvhiIsImdyYWR1YXRlZEZyb20iOiLljZfkuqzoibrmnK/lrabpmaIiLCJpbnRyb2R1Y2UiOiLokZflkI3lm73nlLvlpKfluIgiLCJwaG90b3MiOnsiaGFzaCI6ImY2NDM3NDk2MjY3NzBmMGU0NTkwNjdjY2E2NmEyYmNhMzg4MTVlZjk0Y2JjMzRkYWU5Yzg2YmYxMTI3NjhmYWEiLCJhZGRyZXNzIjoiaHR0cDovLzEwLjEuMy41L3VybC9hZGRyIiwibmFtZSI6IuWktOWDjyJ9LCJjZXJ0aWZpY2F0ZXMiOlt7Imhhc2giOiJhYjI4Yzk0NzdhZDRlYWYyNmUxODAyYjY4ZjM1MmRkODE0MjNiZmZhNDc3NzBlMjJjMzYzN2I0OGM0N2FlZWJiIiwiYWRkcmVzcyI6Imh0dHA6Ly8xMC4xLjMuNS91cmwvZXhhbXBsZSIsIm5hbWUiOiLor4Hku7YxIn1dfQ==","Secret":[]}`

	var t temp
	err := json.Unmarshal([]byte(s), &t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", t)

}
