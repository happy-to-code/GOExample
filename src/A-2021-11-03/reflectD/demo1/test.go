package main

import (
	"fmt"
	"reflect"
)

type Boy struct {
	Name string
	Age  int
}

func main() {
	var b = Boy{"XM", 18}
	fmt.Println("b:", b)
	ty := reflect.TypeOf(b)
	fmt.Println("ty:", ty)
	fmt.Println(reflect.ValueOf(b))
	fmt.Println(reflect.ValueOf(ty))
	fmt.Println("------------------")
	var age interface{} = 25

	fmt.Printf("原始接口变量的类型为 %T，值为 %v \n", age, age)

	t := reflect.TypeOf(age)
	v := reflect.ValueOf(age)

	// 从接口变量到反射对象
	fmt.Printf("从接口变量到反射对象：Type对象的类型为 %T \n", t)
	fmt.Printf("从接口变量到反射对象：Value对象的类型为 %T \n", v)
	fmt.Println(v.Interface())

}
