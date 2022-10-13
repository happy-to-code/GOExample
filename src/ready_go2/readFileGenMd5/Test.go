package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	filePath := "C:\\Users\\user\\Desktop\\go_temp\\demo222.csv"

	t1 := time.Now()
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read fail", err)
	}

	hash := sha256.New()
	hash.Write(f)
	hash256 := hash.Sum(nil)

	toString := hex.EncodeToString(hash256)
	t2 := time.Now()
	fmt.Println("-->", toString)
	fmt.Printf("-->%x\n", hash256)

	fmt.Println(t2.Sub(t1).String())

	// b248bced393b9541566aa3978b65e8aed9b38109f56da7cf798175601e8c81f9
	// 193.6395ms

	// e5ecbd30dc5afa1fda0b01b80893e66771e9648415ba69250b9984dd6da8348f
	// 249.0809ms

	// e5ecbd30dc5afa1fda0b01b80893e66771e9648415ba69250b9984dd6da8348f
}
