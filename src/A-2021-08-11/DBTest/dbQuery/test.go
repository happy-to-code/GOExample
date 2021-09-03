package main

import (
	"flag"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)

func main() {

	var dburl = flag.String("dburl", "", "数据库路径")
	var key = flag.String("key", "", "键")

	flag.Parse()
	// fmt.Println("---->",*dburl,*key)
	// fmt.Println("-2--->",flag.Args())
	db, err := leveldb.OpenFile(*dburl, nil)
	if err != nil {
		panic(fmt.Sprintf("打开数据库失败：%v", err))
	}
	get, err := db.Get([]byte(*key), nil)
	if err != nil {
		panic(fmt.Sprintf("根据键值获取失败：%v", err))
	}
	fmt.Printf("value:%s\n\n", get)

}