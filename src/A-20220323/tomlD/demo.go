package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	FontSize struct {
		LevelOne   int `1.json:"levelOne"`
		LevelTwo   int `1.json:"levelTwo"`
		LevelThree int `1.json:"levelThree"`
		LevelFour  int `1.json:"levelFour"`
	}
	Location map[string][2]int `1.json:"location"`
}

func main() {
	var config Config

	viper.AddConfigPath("E:\\20.06.16Project\\GOExample\\src\\A-20220323\\tomlD")
	viper.SetConfigName("pic")
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Load config file error:", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Unmarshal config error:", err)
	}
	fmt.Printf("%+v\n", config)
}
