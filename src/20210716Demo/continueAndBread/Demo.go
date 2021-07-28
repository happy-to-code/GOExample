package main

import (
	"fmt"
	"sync"
)

type Api struct {
	Name string   `json:"name,omitempty"`
	Urls []string `json:"urls,omitempty"`
}

func main() {
	var e error
	fmt.Println(e == nil)

	var apiList = []Api{
		{
			Name: "Test1",
			Urls: []string{"http://127.0.0.1:80/bbb", "http://127.0.0.1:80/ddd"},
		},
		{
			Name: "Test2",
			Urls: []string{"http://127.0.0.1:90/kkk", "http://127.0.0.1:90/mmm"},
		},
	}

	// var errs []error
	var wg sync.WaitGroup
	var lock sync.Mutex

	var sList []string

	sList = append(sList, "BBBB")

	for _, api := range apiList {
		wg.Add(1)
		fmt.Println(api.Name, "-----------", api.Urls)
		go func(a Api) {
			defer wg.Done()
			for _, url := range a.Urls {
				fmt.Println("url::::", url)
				if url == "http://127.0.0.1:80/bbb" {
					e = fmt.Errorf("errrrrrrrrr")
					lock.Lock()
					fmt.Println("----AAA----")
					lock.Unlock()
					return
				}
				lock.Lock()
				sList = append(sList, "AAAAA")
				lock.Unlock()
			}
		}(api)
	}
	wg.Wait()

	// fmt.Println("err===>", errs)
	fmt.Println("Slist", sList)
	fmt.Println("end")

	fmt.Println(e == nil)
}
