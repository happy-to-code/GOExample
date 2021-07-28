package main

import (
	"fmt"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	// fmt.Printf("topic: %s\n", msg.Topic())
	fmt.Printf("date: %s\n", msg.Payload())
}

type Client struct {
	MQTT.Client
	Url string
}

func SetMqtt(url string) *Client {
	c := &Client{
		Url: url,
	}
	c.initMqtt()
	return c
}
func (m *Client) initMqtt() {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(m.Url)
	m.Client = MQTT.NewClient(opts)
}
func (m *Client) Conn() {
	fmt.Println("conn")

	token := m.Client.Connect()
	if token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
}
func (m *Client) DisConn() {
	fmt.Println("disconn")
	m.Client.Disconnect(250)
}
func (m *Client) Sub() {
	fmt.Println("sub")
	token := m.Client.Subscribe("date", 0, f)
	if token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
}
func (m *Client) Pub() {
	fmt.Println("pub")
	s := time.Now().String()
	token := m.Client.Publish("date", 0, false, s)
	if token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
}
func main() {
	c := SetMqtt("tcp://broker.mqttdashboard.com:1883")
	c.Conn()
	c.Sub()
	c.Pub()
	time.Sleep(time.Second)
	c.DisConn()
}
