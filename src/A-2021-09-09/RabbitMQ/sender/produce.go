package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// 一个辅助函数来检查每个amqp调用的返回值：
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@10.1.3.150:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	for i := 0; i < 1000; i++ {
		body := "1f34gh35f4ghdd:" + fmt.Sprintf("%d", i)
		err = ch.Publish(
			"", // exchange
			// "task_queue", // routing key
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent, // 消息持久化，虽然消息设置持久化了，但是并不能保证一定会
				ContentType:  "text/plain",
				Body:         []byte(body),
			})

		failOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s", body)

	}

}
