package main

import (
	"database/sql"
	"github.com/casbin/casbin/v2"
	casbinpgadapter "github.com/cychiuae/casbin-pg-adapter"
)

func main() {
	// a, err := pgadapter.NewAdapter("postgresql://web_user:web@2022@10.1.4.141:5432/eth_dw?sslmode=disable")

	connectionString := "postgresql://web_user:web890@10.1.4.141:5432/eth_dw?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	tableName := "casbin_rule"
	adapter, err := casbinpgadapter.NewAdapter(db, tableName)
	// If you are using db schema
	// myDBSchema := "mySchema"
	// adapter, err := casbinpgadapter.NewAdapterWithDBSchema(db, myDBSchema, tableName)
	if err != nil {
		panic(err)
	}

	enforcer, err := casbin.NewEnforcer("E:\\20.06.16Project\\GOExample\\src\\A-20220323\\casbin\\rbac_model.conf", adapter)
	if err != nil {
		panic(err)
	}

	// Load stored policy from database
	enforcer.LoadPolicy()

	// Do permission checking
	enforcer.Enforce("alice", "data1", "write")

	// Do some mutations
	enforcer.AddPolicy("alice", "data2", "write")
	enforcer.RemovePolicy("alice", "data1", "write")

	// Persist policy to database
	enforcer.SavePolicy()
}
