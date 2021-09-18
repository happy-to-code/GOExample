package main

import (
	"A-2021-09-09/RabbitMQ2"
)

func main() {
	two := RabbitMQ2.NewRabbitMqRouting("duExchangeName", "two")
	two.ConsumerRouting()
}
