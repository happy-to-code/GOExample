package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"time"
)

var redisPort int

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("E:\\20.06.16Project\\GOExample\\src\\20210716Demo\\viperD")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file:%s Op:%s\n", e.Name, e.Op)
	})
	redisPort = viper.GetInt("redis.port")

	fmt.Println("redis port before sleep: ", redisPort)
	time.Sleep(time.Second * 10)
	fmt.Println("redis port after sleep: ", redisPort)

}
