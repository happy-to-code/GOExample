package main

import (
	"A-2021-09-09/RabbitMQ2"
)

func main() {
	one := RabbitMQ2.NewRabbitMqRouting("duExchangeName", "one")
	one.ConsumerRouting()
}
