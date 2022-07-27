package main

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 配置路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"username": "name1",
			"data":     "data1",
		})
	})
	r.GET("/get", func(c *gin.Context) {
		u := uuid.NewV4().String()
		// c.Writer.Write([]byte(u))
		c.JSON(200, gin.H{
			"uuid": u,
		})
	})
	// 启动 HTTP 服务，默认在 0.0.0.0:8080 启动服务
	r.Run(":8899")
}
