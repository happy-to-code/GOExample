package main

import (
	"A-2021-09-09/RabbitMQ2"
	"fmt"
	"strconv"
)

func main() {
	rabbitmq := RabbitMQ2.NewRabbitMQSimple("duQueueName191224")
	for i := 0; i < 100; i++ {
		rabbitmq.PublishSimple("hello du message" + strconv.Itoa(i) + "---来自work模式")
		fmt.Printf("work模式，共产生了%d条消息\n", i)
	}
}
