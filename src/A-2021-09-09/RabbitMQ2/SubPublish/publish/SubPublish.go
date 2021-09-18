package main

import (
	"A-2021-09-09/RabbitMQ2"
	"fmt"
	"strconv"
)

func main() {
	rabbitmq := RabbitMQ2.NewRabbitMqSubscription("duexchangeName")
	for i := 0; i < 100; i++ {
		rabbitmq.PublishSubscription("订阅模式生产第" + strconv.Itoa(i) + "条数据")
		fmt.Printf("订阅模式生产第" + strconv.Itoa(i) + "条数据\n")
	}
}
