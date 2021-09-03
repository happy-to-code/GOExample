package main

import (
	"fmt"
	levelAdmin "github.com/qjues/leveldb-admin"
	"github.com/syndtr/goleveldb/leveldb"
	"net/http"
	"time"
)

func main() {
	// 其他server 处理
	http.HandleFunc("/other", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello other"))
	})
	go http.ListenAndServe(":4333", nil)

	db, err := leveldb.OpenFile("/root/zhangyifei/peer/peerdb/event.db", nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// 正常的db操作逻辑...
	// 正常的db操作逻辑...
	// err = db.Put([]byte("tizi"), []byte("www.tizi365.com"), nil)
	// err = db.Put([]byte("test"), []byte("demo"), nil)
	// if err != nil {
	// 	panic(err)
	// }

	// 独立端口模式:
	// 只需要向leveldbAdmin注册db的指针和一个用于区分的描述
	// levelAdmin.GetLevelAdmin().Register(db, "description").Start()
	levelAdmin.GetLevelAdmin().Register(db, "description").SetServerMux(http.DefaultServeMux).Start()
	time.Sleep(time.Second * 300)

}
