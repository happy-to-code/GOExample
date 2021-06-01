package main

import (
	"net/http"
	_ "net/http/pprof"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(
	handler appHandler) func(
	http.ResponseWriter, *http.Request) { // 函数作为参数和返回值
	return func(writer http.ResponseWriter,
		request *http.Request) {
		handler(writer, request)

	}
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func main() {
	http.HandleFunc("/", errWrapper(HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
