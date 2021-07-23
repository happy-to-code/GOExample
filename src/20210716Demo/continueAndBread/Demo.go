package main

import "fmt"

type Api struct {
	Name string   `json:"name,omitempty"`
	Urls []string `json:"urls,omitempty"`
}

func main() {
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

	for _, api := range apiList {
		fmt.Println(api.Name, "-----------", api.Urls)
		for _, url := range api.Urls {
			fmt.Println("url::::", url)
			if url == "http://127.0.0.1:80/bbb" {
				break
			}
		}
	}
}
