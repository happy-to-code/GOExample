package main

import "A-2021-09-09/RabbitMQ2"

func main() {
	jay := RabbitMQ2.NewRabbitMqTopic("exchangeNameTpoic1224", "Singer.*")
	jay.ConsumerTopic()
}
