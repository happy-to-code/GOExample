package main

import (
	"A-2021-09-09/RabbitMQ2"
	"fmt"

	"strconv"
)

func main() {
	rabbitmq1 := RabbitMQ2.NewRabbitMqRouting("duExchangeName", "one")
	rabbitmq2 := RabbitMQ2.NewRabbitMqRouting("duExchangeName", "two")
	rabbitmq3 := RabbitMQ2.NewRabbitMqRouting("duExchangeName", "three")
	for i := 0; i < 100; i++ {
		rabbitmq1.PublishRouting("路由模式one" + strconv.Itoa(i))
		rabbitmq2.PublishRouting("路由模式two" + strconv.Itoa(i))
		rabbitmq3.PublishRouting("路由模式three" + strconv.Itoa(i))
		fmt.Printf("在路由模式下，routingKey为one,为two,为three的都分别生产了%d条消息\n", i)
	}
}
