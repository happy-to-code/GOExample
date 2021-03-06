package main

import (
	"fmt"
	"github.com/casbin/casbin"
	pgadapter "github.com/casbin/casbin-pg-adapter"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"net/http"
)

func main() {
	// 要使用自己定义的数据库rbac_db,最后的true很重要.默认为false,使用缺省的数据库名casbin,不存在则创建
	// a := gormadapter.NewAdapter("mysql", "root:root@(127.0.0.1:3306)/casbin_test?charset=utf8&parseTime=True&loc=Local", true)
	// a, err := pgadapter.NewAdapter("postgresql://web_user:web@2022@10.1.4.141:5432/eth_dw?sslmode=disable")
	a, err := pgadapter.NewAdapter(&pg.Options{
		Database: "eth_dw",
		User:     "admin",
		Password: "admin@2022",
		Addr:     fmt.Sprintf("%s:%d", "10.1.4.141", 5432),
	})
	if err != nil {
		fmt.Println("1==>", err)
	}

	e := casbin.NewEnforcer("E:\\20.06.16Project\\GOExample\\src\\A-20220323\\casbin\\rbac_model.conf", a)
	if err != nil {
		fmt.Println("2==>", err)
	}
	// 从DB加载策略
	e.LoadPolicy()

	// 获取router路由对象
	r := gin.New()

	r.POST("/api/v1/add", func(c *gin.Context) {
		fmt.Println("增加Policy")
		// AddPolicy 向当前策略添加授权规则。如果规则已经存在，函数返回false，不会添加规则。否则，该函数通过添加新规则返回 true
		if ok := e.AddPolicy("admin", "/api/v1/hello", "GET"); !ok {
			fmt.Println("Policy已经存在")
		} else {
			fmt.Println("增加成功")
		}
	})
	// 删除policy
	r.DELETE("/api/v1/delete", func(c *gin.Context) {
		fmt.Println("删除Policy")
		// RemovePolicy 从当前策略中删除授权规则。
		if ok := e.RemovePolicy("admin", "/api/v1/hello", "GET"); !ok {
			fmt.Println("Policy不存在")
		} else {
			fmt.Println("删除成功")
		}
	})
	// 获取policy
	r.GET("/api/v1/get", func(c *gin.Context) {
		fmt.Println("查看policy")
		// GetPolicy 获取策略中的所有授权规则。
		list := e.GetPolicy()
		for _, vlist := range list {
			for _, v := range vlist {
				fmt.Printf("value: %s, ", v)
			}
		}
	})
	// 使用自定义拦截器中间件
	r.Use(Authorize(e))
	// 创建请求
	r.GET("/api/v1/hello", func(c *gin.Context) {
		fmt.Println("Hello 接收到GET请求..")
	})

	r.Run(":9000") // 参数为空 默认监听8080端口
}

// 拦截器
func Authorize(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := "admin"

		// 判断策略中是否存在
		if ok := e.Enforce(sub, obj, act); ok {
			fmt.Println("恭喜您,权限验证通过")
			c.Next()
		} else {
			fmt.Println("很遗憾,权限验证没有通过")
			c.Abort()
			c.String(http.StatusUnauthorized, "无权访问")
		}
	}
}
