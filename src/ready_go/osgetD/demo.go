package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func main() {
	os.Setenv("DB_HOME", "/usr/local")
	home := os.Getenv("HOME")
	fmt.Println("home:", home)

	gopath := os.Getenv("GOPATH")
	fmt.Println("gopath:", gopath)

	os.Setenv("DB_USER", "root123")
	viper.SetEnvPrefix("db")
	viper.BindEnv("user")
	envUser := getEnv("user", "etl_develop")
	fmt.Println("envUser:", envUser)
}
func getEnv(key string, defaultValue string) string {
	v := viper.GetString(key)
	if len(v) <= 0 {
		return defaultValue
	}
	return v
}
