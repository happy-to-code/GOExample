package main

import (
	"A-2021-09-09/RabbitMQ2"
	"fmt"
)

func main() {
	rabbitmq := RabbitMQ2.NewRabbitMQSimple("duQueueName1912161843")
	rabbitmq.ConsumeSimple()
	fmt.Println("接收成功！")
}
