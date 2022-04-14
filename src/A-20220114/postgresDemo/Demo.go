package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB
var err error

// PGs数据库信息
const (
	pg_host     = "10.1.3.119"
	pg_port     = 5432
	pg_user     = "postgres"
	pg_password = "wutongchain@gp"
	pg_dbname   = "postgres"
)

func InsertNodeInfo() error {
	stmt, err := db.Prepare("INSERT INTO \"test_user\"(\"id\", \"name\", \"age\") VALUES ($1, $2, $3)")
	if err != nil {
		log.Fatal("PG Statements Wrong: ", err)
	}

	res, err := stmt.Exec(2, "xm", 17)
	if err != nil {
		log.Fatal("PG Statements Exec Wrong: ", err)
	}

	id, err := res.RowsAffected()
	if err != nil {
		log.Fatal("PG Affecte Wrong: ", err)
	}

	fmt.Println(id)
	return nil
}

func BatchInsert() error {
	stmt, err := db.Prepare("INSERT INTO \"test_user\"(\"id\", \"name\", \"age\") VALUES (3, 'xm', 17),(4, 'xh', 18)")
	if err != nil {
		log.Fatal("PG Statements Wrong: ", err)
	}

	res, err := stmt.Exec()
	if err != nil {
		log.Fatal("PG Statements Exec Wrong: ", err)
	}

	id, err := res.RowsAffected()
	if err != nil {
		log.Fatal("PG Affecte Wrong: ", err)
	}

	fmt.Println(id)
	return nil
}

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", pg_host, pg_port, pg_user, pg_password, pg_dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Connect PG Failed: ", err)
	}

	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Ping GP Failed: ", err)
	}
}

func main() {
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", pg_host, pg_port, pg_user, pg_password, pg_dbname)
	// db, err = sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	log.Fatal("Connect PG Failed: ", err)
	// }
	//
	// db.SetMaxOpenConns(2000)
	// db.SetMaxIdleConns(1000)
	// defer db.Close()
	//
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal("Ping GP Failed: ", err)
	// }
	// fmt.Println("PG Successfull Connected!")

	// 插入数据
	err := BatchInsert()
	if err != nil {
		log.Fatal("Insert Data Failed: ", err)
	}

	// 删除数据
	// err := DeleteNodeInfo()
	// if err != nil {
	//	log.Fatal("Delete Data Failed: ", err)
	// }

	// 查询所有数据
	// err := SelectAllNodeInfo()
	// if err != nil {
	//	log.Fatal("Select All Data Failed: ", err)
	// }

	// 更新某一数据
	// err := UpdateNodeInfo()
	// if err != nil {
	// 	log.Fatal("Update Data Failed: ", err)
	// }
}
