package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Event struct {
	Topic       string      `1.json:topic,omitempty`
	SubTopic    string      `1.json:subTopic,omitempty`
	Date        string      `1.json:date,omitempty`        // 事件发生时间
	LedgerID    string      `1.json:ledgerID,omitempty`    // 触发事件的链
	BlockHeight uint64      `1.json:blockHeight,omitempty` // 触发事件的区块
	TxList      []TxInfo    `1.json:txList,omitempty`      // 交易信息
	Client      string      `1.json:client,omitempty`      // 事件发起方——sdk ID
	EventMsg    interface{} `1.json:eventMsg,omitempty`    // 事件消息
	// Dump     bool      //持久化
}
type TxInfo struct {
	TxHash  []byte // 触发事件的交易
	TxIndex uint   // 触发事件的交易索引
}

type Callback struct {
	Ledgerid    string `1.json:ledgerid,omitempty`
	Msgcode     string `1.json:msgcode,omitempty`
	Sender      string `1.json:sender,omitempty`
	Receiver    string `1.json:receiver,omitempty`
	Reftx       string `1.json:reftx,omitempty`
	Msgdata     string `1.json:msgdata,omitempty`
	Cipherkey   string `1.json:cipherkey,omitempty`
	BlockHeight uint64 `1.json:blockHeight,omitempty`
	TxIndex     uint64 `1.json:txIndex,omitempty`
	TxHash      string `1.json:txHash,omitempty`
}

func main() {
	runWeb("8088")
}

func runWeb(port string) {
	gin.SetMode(gin.ReleaseMode)
	// router := gin.Default()
	router := gin.New() // 全局中间件
	// router.Use(gin.Logger()) // 使用 Logger 中间件
	router.Use(gin.Recovery()) // 使用 Recovery 中间件

	router.POST("/event", handlerEvent)
	router.POST("/msg", handlerMsg)
	router.POST("/callback", handlerCallback)
	router.GET("/test/:param", handParam)
	fmt.Println(fmt.Sprintf("服务启动...端口为%v...", port))
	router.Run(":" + port)
}

type MsgToPush struct {
	Content     string `1.json:"content"`
	BlockHeight uint64 `1.json:"blockHeight"`
}

func handlerMsg(c *gin.Context) {
	var req MsgToPush
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err : %s", err))
		return
	}
	fmt.Printf("====>%+v\n", req)
	header := c.GetHeader("msg-topic")
	fmt.Println("header:", header)
}

func handParam(c *gin.Context) {
	param := c.Param("param")

	fmt.Println("param:", param)
	return
}

func handlerEvent(c *gin.Context) {
	var req map[string]interface{}
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err : %s", err))
		return
	}
	eventtmp, err := json.Marshal(req)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Json序列化出错err : %s", err))
		return
	}
	fmt.Printf("%+v\n", string(eventtmp))
	// updateMsg(string(eventtmp), "./eventData.txt")
	c.String(http.StatusOK, "SUCCESS")
}

func handlerCallback(c *gin.Context) {
	var req Callback
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err : %s", err))
		return
	}

	callbacktmp, _ := json.Marshal(req)
	fmt.Printf("%+v\n", string(callbacktmp))
	fmt.Println("")
	updateMsg(string(callbacktmp), "./callBackData.txt")
	c.String(http.StatusOK, "SUCCESS")
}

func updateMsg(msg string, filepath string) {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = file.WriteString(msg)
	}
}
