package Demo1

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_mqtt(t *testing.T) {
	var (
		clientId = "pibigstar"
		wg       sync.WaitGroup
	)
	client := NewClient(clientId)
	err := client.Connect()
	if err != nil {
		t.Errorf(err.Error())
	}

	wg.Add(1)
	go func() {
		err := client.Subscribe(func(c *Client, msg *Message) {
			fmt.Printf("接收到消息: %+v \n", msg)
			wg.Done()
		}, 1, "mqtt")

		if err != nil {
			panic(err)
		}
	}()

	msg := &Message{
		ClientID: clientId,
		Type:     "text",
		Data:     "Hello Pibistar",
		Time:     time.Now().Unix(),
	}
	data, _ := json.Marshal(msg)

	err = client.Publish("mqtt", 1, false, data)
	if err != nil {
		panic(err)
	}

	wg.Wait()
}
