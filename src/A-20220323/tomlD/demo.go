package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	FontSize struct {
		LevelOne   int `json:"levelOne"`
		LevelTwo   int `json:"levelTwo"`
		LevelThree int `json:"levelThree"`
		LevelFour  int `json:"levelFour"`
	}
	Location map[string][2]int `json:"location"`
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
