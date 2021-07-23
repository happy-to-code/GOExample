package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now())
	s := "2c7f6f353d828e99692bb8bf960186f218674581495b399dlb753c00dd636c4f0583f7a833ce67d352e7d32be5d6e3fc899d7004efe1f450fc1a078ee856a8b75"
	fmt.Println(len(s))
	fmt.Println(len("a6080d049c2e4b3585f040b6cbceddfc"))

	timeStr := time.Unix(time.Now().Unix(), 0).Format("06-01-02-15-04-05")
	fmt.Println(timeStr)
}
