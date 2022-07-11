package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	runWeb("8899")
}

func runWeb(port string) {
	gin.SetMode(gin.ReleaseMode)
	// router := gin.Default()
	router := gin.New() // 全局中间件
	// router.Use(gin.Logger()) // 使用 Logger 中间件
	router.Use(gin.Recovery()) // 使用 Recovery 中间件

	router.POST("/event", handlerEvent)
	router.POST("/callback", handlerCallback)
	router.POST("/testParameter", handlerParameter)

	router.GET("/test/:param", handParam)
	fmt.Println(fmt.Sprintf("服务启动...端口为%v...", port))
	router.Run(":" + port)
}

type Boy struct {
	Age  int
	Name string
}

func handlerEvent(c *gin.Context) {
	var boy Boy
	// 参数绑定
	c.BindJSON(&boy)
	fmt.Println("---->", boy)

	c.PostForm(boy.Name)
	handlerCallback(c)
}

type Temp struct {
	Name string
}

func handlerCallback(c *gin.Context) {
	var t Temp
	// 参数绑定
	c.BindJSON(&t)

	fmt.Println("tttttt===>", t)
	t.Name = t.Name + " Change"
	c.JSON(200, t.Name)
}

func handParam(c *gin.Context) {
	param := c.Param("param")

	fmt.Println("param:", param)
	fmt.Printf("param:%T", param)

	c.JSON(200, param+" hello,world")
	return
}

type ParameterTest struct {
	Id         string `1.json:"id,omitempty"`
	DidJsonStr string `1.json:"didJsonStr,omitempty"`
}

func handlerParameter(c *gin.Context) {
	var parameter ParameterTest
	err := c.BindJSON(&parameter)
	if err != nil {
		panic(err)
	}
	fmt.Printf("参数：%+v\n", parameter)

	c.JSON(200, "hello,world")
}
