package main

import (
	"log"
	"os"
)

func main() {
	// 如果文件logs.txt不存在，会自动创建
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	log.Println("Content from log package!")
}
