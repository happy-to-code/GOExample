package main

import "fmt"

type SubType string

type Msg struct {
	ID       uint64  `json:"id"` // 事件的ID
	SubTopic SubType `json:"subTopic"`
}

func main() {
	// s := "宗地面积2217.23㎡/房屋建筑面积135.50㎡"
	//
	// re := regexp.MustCompile("[0-9]+")
	// fmt.Println(re.FindAllString(s, -1))

	// var msg = Msg{
	// 	ID:       123,
	// 	SubTopic: SubType("subtopicTest"),
	// }
	// fmt.Println(msg)
	// fmt.Println(msg.SubTopic)
	// fmt.Println(msg.SubTopic == "subtopicTest")
	var handledBlockNum int64 = 2
	for i := handledBlockNum + 1; ; i++ {
		fmt.Println(i)
		if i == 100 {
			break
		}
	}

}
