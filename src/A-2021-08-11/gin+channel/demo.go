package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// var chanD = make(chan int, 1000000)

func main() {
	fmt.Println("---------------")

	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/post/:name/:password", HinHandler)

	r.Run(":8088") // listen and serve on 0.0.0.0:8080

}

func HinHandler(c *gin.Context) {
	name := c.Param("name")
	pwd := c.Param("password")
	fmt.Println(name, "---", pwd)

	atoi, _ := strconv.Atoi(pwd)
	ed.NumList = append(ed.NumList, atoi)
	if len(ed.Notice) > 0 {
		<-ed.Notice
	}
	ed.Notice <- true

	// 内部提供方法用来返回字符串
	c.String(http.StatusOK, "你好 %s,你的密码为%s", name, pwd)

}

var ed *EventDealer

func init() {
	if ed == nil {
		ed = &EventDealer{
			// NumList: []int{1, 2, 3},
			Notice: make(chan bool, 100),
		}
	}
	go ed.dealEventMsg()
}

func (ed *EventDealer) dealEventMsg() {
	fmt.Println("------启动协程------")
	for {
		select {
		case num := <-ed.Notice:
			fmt.Println("---------->", num)
			for i, v := range ed.NumList {
				fmt.Println("iiiiiiiiiiiiiiiii:", i, "\t", v)
			}
			ed.NumList = ed.NumList[len(ed.NumList):]
		}
	}
}

type EventDealer struct {
	NumList []int
	Notice  chan bool
}
