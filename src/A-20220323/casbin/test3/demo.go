package main

import (
	"fmt"
	pgadapter "github.com/casbin/casbin-pg-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/go-pg/pg/v10"
)

func main() {
	a, err := pgadapter.NewAdapter(&pg.Options{
		Database: "eth_dw",
		User:     "admin",
		Password: "admin@2022",
		Addr:     fmt.Sprintf("%s:%d", "10.1.4.141", 5432),
	})
	if err != nil {
		fmt.Println("1==>", err)
	}
	e, _ := casbin.NewEnforcer("E:\\20.06.16Project\\GOExample\\src\\A-20220323\\casbin\\rbac_model.conf", a)

	e.LoadFilteredPolicy(&pgadapter.Filter{
		P: []string{"", "data1"},
		G: []string{"alice"},
	})
}
