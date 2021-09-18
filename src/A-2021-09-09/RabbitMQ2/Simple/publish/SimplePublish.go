package main

import (
	"A-2021-09-09/RabbitMQ2"
	"fmt"
)

func main() {
	rabbitmq := RabbitMQ2.NewRabbitMQSimple("duQueueName1912161843")
	rabbitmq.PublishSimple("他是客，你是心上人。 ---来自simple模式")
	fmt.Println("发送成功！")
}
