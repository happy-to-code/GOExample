package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func main() {
	db, err := sql.Open("sqlite3", "E:\\20.06.16Project\\GOExample\\src\\A-2021-08-11\\DBTest\\SQLite\\sqLite.db")
	checkErr(err)

	// 插入数据
	stmt, err := db.Prepare("INSERT INTO test(key, value) values(?,?)")
	checkErr(err)

	t1 := time.Now().Unix()
	for i := 0; i < 1000000; i++ {
		k := fmt.Sprintf("key_%d", i)
		v := fmt.Sprintf("value_%d", i)
		stmt.Exec(k, v)

	}
	t2 := time.Now().Unix()
	fmt.Println("time::::", t2-t1)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
