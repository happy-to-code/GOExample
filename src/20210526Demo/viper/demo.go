package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	// 读取yaml文件
	v := viper.New()
	// 设置读取的配置文件
	v.SetConfigName("a")
	// 添加读取的配置文件路径
	v.AddConfigPath("E:\\20.06.16Project\\GOExample\\src\\20210526Demo\\viper")
	// windows环境下为%GOPATH，linux环境下为$GOPATH
	// v.AddConfigPath("$GOPATH/src/")
	// 设置配置文件类型
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
	}

	fmt.Printf(
		`
		TimeStamp:%s
		CompanyInfomation.Name:%s
		CompanyInfomation.Department:%s`,
		v.Get("TimeStamp"),
		v.Get("CompanyInfomation.Name"),
		v.Get("CompanyInfomation.Department"),
	)

	/*
		result:
		TimeStamp:2018-10-18 10:09:23
		CompanyInfomation.Name:Sunny
		CompanyInfomation.Department:[Finance Design Program Sales]
	*/
	fmt.Println()
	fmt.Println("==========================================")
	getStringSlice := v.GetStringSlice("CompanyInfomation.Department")
	for i, s := range getStringSlice {
		fmt.Println(i, "-----", s)
	}
}
