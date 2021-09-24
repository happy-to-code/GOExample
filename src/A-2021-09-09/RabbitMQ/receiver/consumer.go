package main

import (
	"log"

	"github.com/streadway/amqp"
)

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

	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")
	// 为了保证公平分发，不至于其中某个consumer一直处理，而其他不处理
	// err = ch.Qos(
	// 	1,     // prefetch count  在server收到consumer的ACK之前，预取的数量。为1，表示在没收到consumer的ACK之前，只会为其分发一个消息
	// 	0,     // prefetch size 大于0时，表示在收到consumer确认消息之前，将size个字节保留在网络中
	// 	false, // global  true:Qos对同一个connection的所有channel有效； false:Qos对同一个channel上的所有consumer有效
	// )
	// failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack   不进行自动ACK
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	// 其中Auto ack可以设置为true。如果设为true则消费者一接收到就从queue中去除了，如果消费者处理消息中发生意外该消息就丢失了。
	// 如果Auto ack设为false。consumer在处理完消息后，调用msg.Ack(false)后消息才从queue中去除。
	// 即便当前消费者处理该消息发生意外，只要没有执行msg.Ack(false)那该消息就仍然在queue中，不会丢失。
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			// if string(d.Body)=="1f34gh35f4ghdd:999" {
			// 	continue
			// }
			log.Printf("Received a message: %s", d.Body)
			d.Ack(false) // 手动ACK，如果不ACK的话，那么无法保证这个消息被处理，可能它已经丢失了（比如消息队列挂了）
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
