package main

import (
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strings"
)

func createUUID() string {
	ul := uuid.NewV4()
	// 33b7f80e034c426e98ae96bf9dbe7fbe
	// af5991e053584cd1a1fd194de62ad7f2
	return strings.Replace(ul.String(), "-", "", -1)
}
func main() {
	fmt.Println(createUUID())
	s := "[\"http://127.0.0.1:9090/jml/v1/query\",\"http://127.0.0.1:8080/jml/v2/query2\",\"http://127.0.0.1:7070/jml/v3/query3\"]"

	var urlList []string
	err := json.Unmarshal([]byte(s), &urlList)
	if err != nil {
		panic(err)
	}
	// for i, s2 := range urlList {
	// 	fmt.Println(i, "---", s2)
	// }

	fmt.Println("--------------------1")

	// rand.Seed(time.Now().UnixNano())
	// for i := 0; i < 10; i++  {
	// 	println(rand.Intn(3))
	// }

	// rand.Seed(time.Now().UnixNano())
	// for i := 0; i < len(urlList); i++ {
	// 	index := rand.Intn(len(urlList))
	// 	fmt.Println(urlList[index])
	//
	// }

	fmt.Println("--------------------2")

	key2Url := make(map[string]string, len(urlList))
	for _, url := range urlList {
		key2Url[createUUID()] = url
	}
	for k, v := range key2Url {
		fmt.Println(k, "------", v)
	}
}
