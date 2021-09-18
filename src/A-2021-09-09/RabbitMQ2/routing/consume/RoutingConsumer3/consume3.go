package main

import (
	"A-2021-09-09/RabbitMQ2"
)

func main() {
	three := RabbitMQ2.NewRabbitMqRouting("duExchangeName", "three")
	three.ConsumerRouting()
}
