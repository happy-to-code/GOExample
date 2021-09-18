package main

import "A-2021-09-09/RabbitMQ2"

func main() {
	rabbitmq := RabbitMQ2.NewRabbitMQSimple("duQueueName191224")
	rabbitmq.ConsumeSimple()
}
