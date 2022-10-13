package main

import (
	"fmt"
	"github.com/casbin/casbin"
	_ "github.com/go-sql-driver/mysql"
)

func CheckPermi(e *casbin.Enforcer, sub, obj, act string) {
	ok := e.Enforce(sub, obj, act)

	if ok == true {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}
func main() {
	e := casbin.NewEnforcer("E:\\20.06.16Project\\GOExample\\src\\ready_go2\\casbin_d\\model.conf",
		"E:\\20.06.16Project\\GOExample\\src\\ready_go2\\casbin_d\\policy.csv")

	// 基本权限设置
	CheckPermi(e, "demo", "/user", "read")
	CheckPermi(e, "demo", "/order", "write")
	CheckPermi(e, "demo1", "/user/userlist", "read")
	CheckPermi(e, "demo1", "/order/orderlist", "write")
}
