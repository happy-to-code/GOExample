package main

import (
	"fmt"
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// PGs数据库信息
const (
	pg_addr     = "10.1.3.117:5432"
	pg_user     = "gpadmin"
	pg_password = "wutongchain@gp"
	pg_dbname   = "postgres"
)

// 定义表结构
type NodeInfo struct {
	NodeName     string `sql:"type:varchar(45),unique,notnull,pk"`
	NodeIp       string `sql:"type:varchar(45),notnull"`
	NodePort     string `sql:"type:varchar(45),notnull"`
	NodeUsername string `sql:"type:varchar(45),notnull"`
	NodePassword string `sql:"type:varchar(255),notnull"`
}

func CreateTable(db *pg.DB) error {
	for _, model := range []interface{}{(*NodeInfo)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteTable(db *pg.DB) error {

	err := db.DropTable(&NodeInfo{}, &orm.DropTableOptions{
		IfExists: true,
		Cascade:  true,
	})

	return err
}

func InsertNodeInfo(db *pg.DB) error {
	// 插入数据方法一
	// nodeinfodata := &NodeInfo{
	//	NodeName:     "chenjian",
	//	NodeIp:       "10.0.0.5",
	//	NodePort:     "2222",
	//	NodeUsername: "chenjian",
	//	NodePassword: "1234321",
	// }
	// err := db.Insert(nodeinfodata)
	// if err!=nil {
	//	return err
	// }

	// 插入数据方法二
	err := db.Insert(&NodeInfo{
		NodeName:     "chenjian1",
		NodeIp:       "10.0.0.5",
		NodePort:     "2222",
		NodeUsername: "chenjian",
		NodePassword: "1234321",
	})
	if err != nil {
		return err
	}
	return nil
}

func SelectAllNodeInfo(db *pg.DB) error {
	var nodeinfodata []NodeInfo
	err := db.Model(&nodeinfodata).Select()
	if err != nil {
		return err
	}
	fmt.Println(nodeinfodata)
	return nil
}

func SelectNodeInfo(db *pg.DB) error {
	nodeinfodata := &NodeInfo{
		NodeName: "chenjian",
	}
	err := db.Select(nodeinfodata)
	if err != nil {
		return err
	}
	fmt.Println(nodeinfodata)
	return nil
}

func UpdateNodeInfo(db *pg.DB) error {
	var nodeinfodata []NodeInfo

	updatedata := &NodeInfo{
		NodeName:     "chenjian2",
		NodeIp:       "10.0.0.5",
		NodePort:     "3333",
		NodeUsername: "chenjian",
		NodePassword: "1234321",
	}

	_, err := db.Model(&nodeinfodata).
		Set("node_port = ?", updatedata.NodePort).
		Where("node_name = ?", updatedata.NodeName).
		Returning("*").
		Update()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// 链接PostgreSQL数据库
	log.Println("Connecting PostgreSQL....")

	db := pg.Connect(&pg.Options{
		Addr:     pg_addr,
		User:     pg_user,
		Password: pg_password,
		Database: pg_dbname,
	})
	defer db.Close()

	fmt.Println("Successfull Connected!")

	// 创建表
	// err := CreateTable(db)
	// if err != nil {
	// 	log.Fatal("Create Table Failed: ", err)
	// }

	// 删除表
	// err := DeleteTable(db)
	// if err != nil {
	//	log.Fatal("Delete Table Failed: ", err)
	// }

	// 插入数据
	err := InsertNodeInfo(db)
	if err != nil {
		log.Fatal("Insert Data Failed: ", err)
	}

	// 查询所有数据
	// err := SelectAllNodeInfo(db)
	// if err != nil {
	//	log.Fatal("Select All Data Failed: ", err)
	// }

	// 查询某一数据
	// err := SelectNodeInfo(db)
	// if err != nil {
	//	log.Fatal("Select Data Failed: ", err)
	// }

	// 更新某一数据
	// err := UpdateNodeInfo(db)
	// if err != nil {
	// 	log.Fatal("Update Data Failed: ", err)
	// }
}
