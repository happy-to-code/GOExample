package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

var s = `{}`

func main() {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		panic(err)
	}
	fmt.Println("-->", m)

	// var path = "product_information.essential_information.basic_product_information.product_issuer_subject_ref"
	// productIssuerSubjectRef := gjson.Get(s, path)
	//
	// fmt.Println("productIssuerSubjectRef:", productIssuerSubjectRef)
	// // println(productIssuerSubjectRef.Type)
	// fmt.Println("===>", productIssuerSubjectRef.Str+"/"+"11")
	// set, err := sjson.Set(s, path, productIssuerSubjectRef.Str+"/"+"11")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(set)

	// var path2 = "product_information.test[0].mytest_ref"
	var path2 = "product_information.test"
	get := gjson.Get(s, path2)
	fmt.Println("------>", get)
	fmt.Println("------>", get.Type)
	fmt.Println("------>", get.IsArray())
	result := gjson.Get(s, path2+".0.mytest_ref")
	fmt.Println("result:", result)
	set, err := sjson.Set(s, "product_information.test.0.mytest_ref", "cvb/123")
	if err != nil {
		panic(err)
	}
	fmt.Println("====", set)

	// fmt.Println("========================================================================")
	// var path = "product_information.essential_information.basic_product_information.product_issuer_subject_ref"
	// productIssuerSubjectRef := gjson.Get(s, path)
	// fmt.Println(productIssuerSubjectRef.Str)
	// fmt.Println(productIssuerSubjectRef.IsArray())
}
