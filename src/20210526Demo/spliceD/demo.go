package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var dis []DataItem
	dis = append(dis, DataItem{"1", "婚姻", "http://aaa", 0, ""})
	dis = append(dis, DataItem{"2", "社保", "http://bbb", 0, "2021"})

	for i, item := range dis {
		if strings.TrimSpace(item.CreateTime) == "" {
			item.CreateTime = time.Now().Format("2006-01-02 15:04:05")
			dis[i] = item
		}
	}
	fmt.Println(dis)

}

type DataItem struct {
	Id string `1.json:"id" binding:"required"` //  唯一ID
	// AuthId      string `1.json:"authId"`                  //  接口拥有者
	Name        string `1.json:"name" binding:"required"` //  接口名称
	Url         string `1.json:"url" binding:"required"`  //  接口路径
	RequestType int    `1.json:"requestType"`             //  请求类型(0:GET;1:POST),默认是0
	CreateTime  string `1.json:"createTime"`              //  创建时间（格式：yyyy-MM-dd HH:mm:ss），默认是当前时间
	// UseType     int    `1.json:"useType"`                 //  类型（0：自然人；1：法人；2：通用）默认是0
}
