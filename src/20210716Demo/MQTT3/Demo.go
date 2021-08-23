package main

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// 创建全局mqtt publish消息处理 handler
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Println("发布消息：")
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

// 创建全局mqtt sub消息处理 handler
var messageSubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Println("收到订阅消息：")
	fmt.Printf("Sub Client Topic : %s \n", msg.Topic())
	fmt.Printf("Sub Client msg : %s \n", msg.Payload())
}

func main() {
	// 创建消息发布go程
	go func() {
		opts := mqtt.NewClientOptions().AddBroker("tcp://10.1.3.151:1884").SetClientID("emqx_test_client66666")
		opts.SetKeepAlive(60 * time.Second)
		// Message callback handler，在没有任何订阅时，发布端调用此函数
		opts.SetDefaultPublishHandler(messagePubHandler)
		opts.SetPingTimeout(1 * time.Second)

		client := mqtt.NewClient(opts)
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

		for i := 0; i < 6; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println("发布一条天气预报消息")
			token := client.Publish("mytesttopic/1", 0, false, "Hello World，苏州天气晴朗")
			token.Wait()
		}

		fmt.Println("发布客户端断开与broker的连接")
		client.Disconnect(250)
	}()

	// 创建消息订阅go程
	go func() {
		opts := mqtt.NewClientOptions().AddBroker("tcp://10.1.3.151:1884").SetClientID("emqx_test_client99999")
		opts.SetKeepAlive(60 * time.Second)
		opts.SetPingTimeout(1 * time.Second)

		client := mqtt.NewClient(opts)
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

		// Subscription
		/*if token := client.Subscribe("testtopic/#", 0, messageSubHandler); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}*/
		for i := 0; i < 6; i++ {
			time.Sleep(time.Duration(3) * time.Second)
			// 发布消息
			if token := client.Subscribe("mytesttopic/#", 1, messageSubHandler); token.Wait() && token.Error() != nil {
				fmt.Println(token.Error())
				os.Exit(1)
			}
		}

		// Cancel subscription
		// if token := client.Unsubscribe("mytesttopic/#"); token.Wait() && token.Error() != nil {
		// 	fmt.Println(token.Error())
		// 	os.Exit(1)
		// }

		fmt.Println("订阅客户端断开与broker的连接")
		// c.Disconnect(250)
	}()

	// Disconnect
	time.Sleep(10 * time.Second)
}
