package main

import (
	"fmt"
	levelAdmin "github.com/qjues/leveldb-admin"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"net/http"
	"time"
)

func main() {
	// 其他server 处理
	http.HandleFunc("/other", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello other"))
	})
	go http.ListenAndServe(":4333", nil)

	db, err := leveldb.OpenFile("E:\\20.06.16Project\\GOExample\\src\\A-2021-08-11\\DBTest\\levelDB\\leveldb", nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// 正常的db操作逻辑...
	// 正常的db操作逻辑...
	err = db.Put([]byte("key_1"), []byte("key_1"), nil)
	err = db.Put([]byte("key_2"), []byte("key_2"), nil)
	err = db.Put([]byte("key_3"), []byte("key_3"), nil)
	err = db.Put([]byte("test"), []byte("test"), nil)

	// 独立端口模式:
	// 只需要向leveldbAdmin注册db的指针和一个用于区分的描述
	// levelAdmin.GetLevelAdmin().Register(db, "description").Start()
	levelAdmin.GetLevelAdmin().Register(db, "description").SetServerMux(http.DefaultServeMux).Start()

	iter := db.NewIterator(util.BytesPrefix([]byte("key_")), nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Println(string(key), "=>", string(value))
	}
	iter.Release()
	iter.Error()

	time.Sleep(time.Second * 300)

}
