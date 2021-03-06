package main

import (
	"errors"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type User1 struct {
	Id     int64
	Name   string
	Emails []string
}

func (u User1) String() string {
	return fmt.Sprintf("User1<%d %s %v>", u.Id, u.Name, u.Emails)
}

func main() {
	var (
		err         error
		pgsqlDB     *pg.DB
		result      pg.Result
		user1       *User1
		updateUser  User1
		delUser     User1
		userList    []User1
		queryResult []User1
	)

	// 1.连接数据库
	pgsqlDB = pg.Connect(&pg.Options{
		Addr:     "121.239.164.179:22021",
		User:     "root",
		Password: "wutongchain@ps",
		Database: "testdb",
	})
	if pgsqlDB == nil {
		err = errors.New("pg.Connect() failed,error:")
		goto ERR
	}

	// 莫忘记关闭数据库连接
	defer func(pgsqlDB *pg.DB) {
		err := pgsqlDB.Close()
		if err != nil {
			fmt.Println("close postgresql failed")
		}
	}(pgsqlDB)

	// 3.创建表
	err = createSchema(pgsqlDB)
	if err != nil {
		goto ERR
	}

	// 4.插入一条记录
	user1 = &User1{
		Id:     1,
		Name:   "admin",
		Emails: []string{"admin1@admin", "admin2@admin"},
	}
	result, err = pgsqlDB.Model(user1).Insert()
	if err != nil {
		goto ERR
	}
	fmt.Printf("single insert rows affected:%d\n", result.RowsAffected())

	// 5.批量插入多条记录
	userList = []User1{
		{
			Id:     2,
			Name:   "haolipeng",
			Emails: []string{"1078285863@qq.com"},
		},
		{
			Id:     3,
			Name:   "haolipeng",
			Emails: []string{"haolipeng12345@163.com"},
		},
	}
	result, err = pgsqlDB.Model(&userList).Insert()
	if err != nil {
		goto ERR
	}
	fmt.Printf("batch insert rows affected:%d\n", result.RowsAffected())

	// 6.查询
	err = pgsqlDB.Model(&queryResult).Select()
	if err != nil {
		goto ERR
	}
	fmt.Printf("query result:%v\n", queryResult)

	// 7.修改
	// 修改除主键外的其他列
	updateUser = User1{
		Id:     1,
		Name:   "antiy",
		Emails: []string{"haolipeng@antiy.cn"},
	}
	result, err = pgsqlDB.Model(&updateUser).WherePK().Update()
	if err != nil {
		goto ERR
	}
	fmt.Printf("update rows affected:%d\n", result.RowsAffected())

	// 8.删除记录(删除id为2的记录)
	delUser = User1{
		Id: 2,
	}
	result, err = pgsqlDB.Model(&delUser).WherePK().Delete()
	if err != nil {
		goto ERR
	}
	fmt.Printf("delete rows affected:%d\n", result.RowsAffected())

	// 9.将当前记录查询并都打印出来
	err = pgsqlDB.Model(&queryResult).Select()
	if err != nil {
		goto ERR
	}
	fmt.Printf("query result:%v\n", queryResult)
	return
ERR:
	fmt.Println("error:", err)
	return
}

// 通过结构体来删除表
func deleteSchema(db *pg.DB) error {
	models := []interface{}{
		(*User1)(nil),
	}
	err := db.Model(&models).DropTable(&orm.DropTableOptions{
		IfExists: true,
		Cascade:  true,
	})
	return err
}

// 通过定义的结构体来创建数据库表
func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*User1)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			// Temp: true,//建表是临时的
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
