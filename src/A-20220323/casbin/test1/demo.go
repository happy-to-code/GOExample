package main

import (
	"fmt"
	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/http"
)

// Response 统一响应结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var O *gorm.DB
var PO *gormadapter.Adapter
var Enforcer *casbin.Enforcer

func ping(c *gin.Context) {
	var response Response
	response.Code = 0
	response.Message = "success"
	response.Data = ""
	c.JSON(200, response)
	return
}

// 数据库连接及角色规则的初始化
func connect() {
	dsn := "root:root@(127.0.0.1:3306)/casbin_test?charset=utf8&parseTime=True&loc=Local"
	var err error
	O, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("connect DB error")
		panic(err)
	}
	// 将数据库连接同步给插件， 插件用来操作数据库
	PO = gormadapter.NewAdapterByDB(O)
	// 这里也可以使用原生字符串方式
	Enforcer = casbin.NewEnforcer("E:\\20.06.16Project\\GOExample\\src\\A-20220323\\casbin\\rbac_model.conf", PO)
	// 开启权限认证日志
	Enforcer.EnableLog(true)
	// 加载数据库中的策略
	err = Enforcer.LoadPolicy()
	if err != nil {
		fmt.Println("loadPolicy error")
		panic(err)
	}
	// 创建一个角色,并赋于权限
	// admin 这个角色可以用 GET 方式访问 /api/v2/ping
	res := Enforcer.AddPolicy("admin", "/api/v2/ping", "GET")
	if !res {
		fmt.Println("policy is exist")
	} else {
		fmt.Println("policy is not exist, adding")
	}
	// 将 test 用户加入一个角色中
	Enforcer.AddRoleForUser("test", "root")
	Enforcer.AddRoleForUser("tom", "admin")
	// 请看规则中如果用户名为 root 则不受限制
}

func main() {
	defer O.Close()
	connect()
	g := gin.Default()
	// 这里的接口没有使用权限认证中间件
	version1 := g.Group("/api/v1")
	{
		version1.GET("/ping", ping) // 这个是通用的接口
	}
	// 接口使用权限认证中间件
	version2 := g.Group("/api/v2", CasbinMiddleWare)
	{
		version2.GET("/ping", ping)
	}
	_ = g.Run(":8099")
}

// CasbinMiddleWare casbin middleware 权限认证中间件
func CasbinMiddleWare(c *gin.Context) {
	var userName string
	userName = c.GetHeader("userName")
	if userName == "" {
		fmt.Println("headers invalid")
		c.JSON(200, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Unauthorized",
			"data":    "",
		})
		c.Abort()
		return
	}
	// 请求的path
	p := c.Request.URL.Path
	// 请求的方法
	m := c.Request.Method
	// 这里认证
	res, err := Enforcer.EnforceSafe(userName, p, m)
	// 这个 HasPermissionForUser 跟上面的有什么区别
	// EnforceSafe 会验证角色的相关的权限
	// 而 HasPermissionForUser 只验证用户是否有权限
	// res = Enforcer.HasPermissionForUser(userName,p,m)
	if err != nil {
		fmt.Println("no permission ")
		fmt.Println(err)
		c.JSON(200, gin.H{
			"code":    401,
			"message": "Unauthorized",
			"data":    "",
		})
		c.Abort()
		return
	}
	if !res {
		fmt.Println("permission check failed")
		c.JSON(200, gin.H{
			"code":    401,
			"message": "Unauthorized",
			"data":    "",
		})
		c.Abort()
		return
	}
	c.Next()
}
