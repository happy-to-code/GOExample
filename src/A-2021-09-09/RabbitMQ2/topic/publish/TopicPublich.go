package main

import (
	"A-2021-09-09/RabbitMQ2"
	"fmt"
	"strconv"
)

func main() {
	one := RabbitMQ2.NewRabbitMqTopic("exchangeNameTpoic1224", "Singer.Jay")
	two := RabbitMQ2.NewRabbitMqTopic("exchangeNameTpoic1224", "Persident.XIDADA")
	for i := 0; i < 100; i++ {
		one.PublishTopic("小杜同学，topic模式，Jay," + strconv.Itoa(i))
		two.PublishTopic("小杜同学，topic模式，All," + strconv.Itoa(i))
		fmt.Printf("topic模式。这是小杜同学发布的消息%v \n", i)
	}
}
