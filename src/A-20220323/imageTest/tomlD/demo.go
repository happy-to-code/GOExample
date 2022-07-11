package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	picViper := viper.New()
	picViper.AddConfigPath("./conf/middleware")
	picViper.AddConfigPath("E:\\20.06.16Project\\GOExample\\src\\A-20220323\\imageTest\\tomlD\\")
	picViper.SetConfigName("1")
	picViper.SetConfigType("toml")
	if err := picViper.ReadInConfig(); err != nil {
		panic(err)
	}
	picViper.WatchConfig()

	fmt.Println("---------------->")
	fmt.Println(picViper.GetStringMap("StoreCert.Title"))
	fmt.Println(picViper.GetString("StoreCert.Title.name"))
	fmt.Println(picViper.GetIntSlice("StoreCert.Title.location"))

}
