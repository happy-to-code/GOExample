package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strings"
)

func createUUID() string {
	ul := uuid.NewV4()
	return strings.Replace(ul.String(), "-", "", -1)
}
func main() {
	ss := "3130303131393230323130333031313239323933"

	fmt.Println(ss[:8])
	fmt.Println(len(ss))
	for i := 0; i < 100; i++ {
		s := createUUID()
		fmt.Println(s, "--", len(s))
	}
}
