package main

import (
	"fmt"
	"time"
)

func main() {
	time := time.Now().Add(-time.Second * 300)
	fmt.Println(time)
	format := time.Format("2006-01-02 15:04:05")
	fmt.Println(`format: ${format}`, format)
}
