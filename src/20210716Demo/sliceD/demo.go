package main

import "fmt"

func main() {
	topicList := []string{"scevent", "txevent"}

	topic := make([]string, 0)

	for k := range topicList {
		tmp := topicList[k] + "/" + "#"
		topic = append(topic, tmp)
	}
	fmt.Println(topic)
	topic = append(topic, fmt.Sprintf("ReceiveHistory==once;%v", 10))
	fmt.Printf("%+v\n", topic)

	tq := make([]TopicQos, len(topic))
	for i := 0; i < len(topic); i++ {
		tq[i].Topic = topic[i]
		tq[i].Qos = 1
	}

	fmt.Println("-->", tq)
}

type TopicQos struct {
	Topic string
	Qos   QosLevel
}
type QosLevel uint8
