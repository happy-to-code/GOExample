package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int    `1.json:"id"`
	UserName string `1.json:"userName" `
	Password string `1.json:"password"`
}

func main() {
	id := SelectUserById(2)
	fmt.Println(id)
	users := SelectAllUser()
	fmt.Println(users)
}

var DB *sql.DB
var err error

func init() {
	DB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("1 open database fail:", err)
		return
	}
	// 设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	// 设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	// 验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("2 open database fail:", err)
		return
	}
	fmt.Println("connnect success")
}

func SelectUserById(id int) User {
	var user User
	err := DB.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.Id, &user.UserName, &user.Password)
	if err != nil {
		fmt.Println("查询出错了")
	}
	return user
}

func SelectAllUser() (users []User) {
	// 执行查询语句
	rows, err := DB.Query("SELECT * from users")
	if err != nil {
		fmt.Println("查询出错了")
	}
	// var users []User
	// 循环读取结果
	for rows.Next() {
		var user User
		// 将每一行的结果都赋值到一个user对象中
		err := rows.Scan(&user.Id, &user.UserName, &user.Password)
		if err != nil {
			fmt.Println("rows fail")
			return
		}
		// 将user追加到users的这个数组中
		users = append(users, user)
	}
	return users
}
