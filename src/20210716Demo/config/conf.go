package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"time"
)

type Reader interface {
	GetAllKeys() []string
	Get(key string) interface{}
	GetBool(key string) bool
	GetString(key string) string
}
type viperConfigReader struct {
	viper *viper.Viper
}

var TestConfReader *viperConfigReader

func (v viperConfigReader) GetAllKeys() []string {
	return v.viper.AllKeys()
}
func (v viperConfigReader) Get(key string) interface{} {
	return v.viper.Get(key)
}
func (v viperConfigReader) GetBool(key string) bool {
	return v.viper.GetBool(key)
}
func (v viperConfigReader) GetString(key string) string {
	return v.viper.GetString(key)
}

func init() {
	v := viper.New()
	v.SetConfigName("test")
	v.AddConfigPath("/tmp/")
	err := v.ReadInConfig()
	if err != nil {
		log.Panic("Not able to read configuration", err.Error())
	}
	TestConfReader = &viperConfigReader{
		viper: v,
	}
	go func() {
		for {
			time.Sleep(time.Second * 5)
			v.WatchConfig()
			v.OnConfigChange(func(e fsnotify.Event) {
				log.Println("config file changed", e.Name)
			})
		}
	}()
}

func main() {
	conf := TestConfReader
	log.Println(conf.GetAllKeys())
	log.Println(conf.GetString("user"))
	time.Sleep(20 * time.Second)
	log.Println(conf.GetString("user"))
}
