package RabbitMQ2

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// 这里主要是RabbitMQ的一些信息。包括其结构体和函数。

// MQURL 连接信息
const MQURL = "amqp://admin:admin@10.1.3.150:5672/"

// RabbitMQ RabbitMQ结构体
type RabbitMQ struct {
	// 连接
	conn    *amqp.Connection
	channel *amqp.Channel
	// 队列
	QueueName string
	// 交换机名称
	ExChange string
	// 绑定的key名称
	Key string
	// 连接的信息，上面已经定义好了
	MqUrl string
}

// NewRabbitMQ 创建结构体实例，参数队列名称、交换机名称和bind的key（也就是几个大写的，除去定义好的常量信息）
func NewRabbitMQ(queueName string, exChange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, ExChange: exChange, Key: key, MqUrl: MQURL}
}

// Destroy 关闭conn和channel的方法
func (r *RabbitMQ) Destroy() {
	r.channel.Close()
	r.conn.Close()
}

// 错误的函数处理
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		fmt.Printf("err是:%s,小杜同学手写的信息是:%s", err, message)
	}
}

// 01,这里是rabbitmq最简单的模式：simple模式。
// 也就是由生产者将消息送到队列里，然后由消费者到队列里取出来消费。

// NewRabbitMQSimple 创建简单模式下的实例，只需要queueName这个参数，其中exchange是默认的，key则不需要。
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	rabbitmq := NewRabbitMQ(queueName, "", "")
	var err error
	// 获取参数connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)
	rabbitmq.failOnErr(err, "连接connection失败")
	// 获取channel参数
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel参数失败")
	return rabbitmq
}

// PublishSimple 直接模式,生产者.
func (r *RabbitMQ) PublishSimple(message string) {
	// 第一步，申请队列，如不存在，则自动创建之，存在，则路过。
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("创建连接队列失败：%s", err)
	}

	// 第二步，发送消息到队列中
	r.channel.Publish(
		r.ExChange,
		r.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// ConsumeSimple 直接模式，消费者
func (r *RabbitMQ) ConsumeSimple() {
	// 第一步,申请队列,如果队列不存在则自动创建,存在则跳过
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		// 是否持久化
		false,
		// 是否自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞处理
		false,
		// 额外的属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	// 第二步,接收消息
	msgs, err := r.channel.Consume(
		q.Name,
		"",   // 用来区分多个消费者
		true, // 是否自动应答,告诉我已经消费完了
		false,
		false, // 若设置为true,则表示为不能将同一个connection中发送的消息传递给这个connection中的消费者.
		false, // 消费队列是否设计阻塞
		nil,
	)
	if err != nil {
		fmt.Printf("消费者接收消息出现问题:%s", err)
	}

	forever := make(chan bool)
	// 启用协程处理消息
	go func() {
		for d := range msgs {
			log.Printf("Simple模式接收到了消息:%s\n", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

// 这里是订阅模式的相关代码。
// 订阅模式需要用到exchange。

// NewRabbitMqSubscription 获取订阅模式下的rabbitmq的实例
func NewRabbitMqSubscription(exchangeName string) *RabbitMQ {
	// 创建rabbitmq实例
	rabbitmq := NewRabbitMQ("", exchangeName, "")
	var err error
	// 获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)
	rabbitmq.failOnErr(err, "订阅模式连接rabbitmq失败。")
	// 获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "订阅模式获取channel失败")
	return rabbitmq
}

// PublishSubscription 订阅模式发布消息
func (r *RabbitMQ) PublishSubscription(message string) {
	// 第一步，尝试连接交换机
	err := r.channel.ExchangeDeclare(
		r.ExChange,
		"fanout", // 这里一定要设计为"fanout"也就是广播类型。
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "订阅模式发布方法中尝试连接交换机失败。")

	// 第二步，发送消息
	err = r.channel.Publish(
		r.ExChange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// ConsumeSbuscription 订阅模式消费者
func (r *RabbitMQ) ConsumeSbuscription() {
	// 第一步，试探性创建交换机exchange
	err := r.channel.ExchangeDeclare(
		r.ExChange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "订阅模式消费方法中创建交换机失败。")

	// 第二步，试探性创建队列queue
	q, err := r.channel.QueueDeclare(
		"", // 随机生产队列名称
		false,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "订阅模式消费方法中创建创建队列失败。")

	// 第三步，绑定队列到交换机中
	err = r.channel.QueueBind(
		q.Name,
		"", // 在pub/sub模式下key要为空
		r.ExChange,
		false,
		nil,
	)

	// 第四步，消费消息
	messages, err := r.channel.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range messages {
			fmt.Printf("发布订阅模式收到的消息：%s\n", d.Body)
			d.Ack(false)
		}
	}()

	fmt.Println("订阅模式消费者已开启，退出请按 CTRL+C\n")
	<-forever

}

// rabbitmq的路由模式。
// 主要特点不仅一个消息可以被多个消费者消费还可以由生产端指定消费者。
// 这里相对比订阅模式就多了一个routingkey的设计，也是通过这个来指定消费者的。
// 创建exchange的kind需要是"direct",不然就不是roting模式了。

// NewRabbitMqRouting 创建rabbitmq实例，这里有了routingkey为参数了。
func NewRabbitMqRouting(exchangeName string, routingKey string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchangeName, routingKey)
	var err error
	// 获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)
	rabbitmq.failOnErr(err, "创建rabbit的路由实例的时候连接出现问题")
	// 获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "创建rabbitmq的路由实例时获取channel出错")
	return rabbitmq
}

// PublishRouting 路由模式，产生消息。
func (r *RabbitMQ) PublishRouting(message string) {
	// 第一步，尝试创建交换机，与pub/sub模式不同的是这里的kind需要是direct
	err := r.channel.ExchangeDeclare(r.ExChange, "direct", true, false, false, false, nil)
	r.failOnErr(err, "路由模式，尝试创建交换机失败")
	// 第二步，发送消息
	err = r.channel.Publish(
		r.ExChange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// ConsumerRouting 路由模式，消费消息。
func (r *RabbitMQ) ConsumerRouting() {
	// 第一步，尝试创建交换机，注意这里的交换机类型与发布订阅模式不同，这里的是direct
	err := r.channel.ExchangeDeclare(
		r.ExChange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "路由模式，创建交换机失败。")

	// 第二步，尝试创建队列,注意这里队列名称不用写，这样就会随机产生队列名称
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "路由模式，创建队列失败。")

	// 第三步，绑定队列到exchange中
	err = r.channel.QueueBind(q.Name, r.Key, r.ExChange, false, nil)

	// 第四步，消费消息。
	messages, err := r.channel.Consume(q.Name, "", true, false, false, false, nil)
	forever := make(chan bool)
	go func() {
		for d := range messages {
			log.Printf("小杜同学写的路由模式(routing模式)收到消息为：%s。\n", d.Body)
		}
	}()
	<-forever
}

// topic模式
// 与routing模式不同的是这个exchange的kind是"topic"类型的。
// topic模式的特别是可以以通配符的形式来指定与之匹配的消费者。
// "*"表示匹配一个单词。“#”表示匹配多个单词，亦可以是0个。

// NewRabbitMqTopic 创建rabbitmq实例
func NewRabbitMqTopic(exchangeName string, routingKey string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchangeName, routingKey)
	var err error
	// 获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)
	rabbitmq.failOnErr(err, "创建rabbit的topic模式时候连接出现问题")
	// 获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "创建rabbitmq的topic实例时获取channel出错")
	return rabbitmq
}

// PublishTopic topic模式。生产者。
func (r *RabbitMQ) PublishTopic(message string) {
	// 第一步，尝试创建交换机,这里的kind的类型要改为topic
	err := r.channel.ExchangeDeclare(
		r.ExChange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "topic模式尝试创建exchange失败。")

	// 第二步，发送消息。
	err = r.channel.Publish(
		r.ExChange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// ConsumerTopic topic模式。消费者。"*"表示匹配一个单词。“#”表示匹配多个单词，亦可以是0个。
func (r *RabbitMQ) ConsumerTopic() {
	// 第一步，创建交换机。这里的kind需要是“topic”类型。
	err := r.channel.ExchangeDeclare(
		r.ExChange,
		"topic",
		true, // 这里需要是true
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "topic模式，消费者创建exchange失败。")

	// 第二步，创建队列。这里不用写队列名称。
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "topic模式，消费者创建queue失败。")

	// 第三步，将队列绑定到交换机里。
	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.ExChange,
		false,
		nil,
	)

	// 第四步，消费消息。
	messages, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range messages {
			log.Printf("小杜同学写的topic模式收到了消息：%s。\n", d.Body)
		}
	}()
	<-forever
}
