package main

import (
	"net/http"
	_ "net/http/pprof"
)

func funcA() []byte {
	a := make([]byte, 10*1024*1024)
	return a
}

func funcB() ([]byte, []byte) {
	a := make([]byte, 10*1024*1024)
	b := funcA()
	return a, b
}

func funcC() ([]byte, []byte, []byte) {
	a := make([]byte, 10*1024*1024)
	b, c := funcB()
	return a, b, c
}

func main() {
	for i := 0; i < 5; i++ {
		funcA()
		funcB()
		funcC()
	}

	http.ListenAndServe("0.0.0.0:9999", nil)
	// http://127.0.0.1:9999/debug/pprof/
	// go tool pprof http://localhost:9999/debug/pprof/heap
	// go tool pprof http://localhost:9999/debug/pprof/profile
}
