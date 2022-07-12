package main

import (
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

// Task 任务
type Task struct {
}

var taskMap = make(map[string]func())
var d *Demo

func init() {
	taskMap["procTask"] = ProcTask
	taskMap["procTask2"] = d.ProcTask2
}

func ProcTask() {
	log.Println("hello world")
}

type Demo struct {
	Name string
}

func (d *Demo) ProcTask2() {
	log.Println("----hello world----")
}

// Init 初始化
func Init() (obj *Task) {
	obj = &Task{}
	obj.Execute()
	return
}

// Execute 执行任务
func (obj *Task) Execute() {
	crontabList := make(map[string]string)
	// 每分钟执行一次
	// crontabList["procTask"] = "0 */1 * * * *"
	crontabList["procTask"] = "*/1 * * * * ?"  // 每秒执行一次
	crontabList["procTask2"] = "*/2 * * * * ?" // 每2秒执行一次
	c := cron.New(cron.WithSeconds())

	log.Println(crontabList)
	for key, value := range crontabList {
		if _, ok := taskMap[key]; ok {
			_, err := c.AddFunc(value, taskMap[key])
			if err != nil {
				log.Println("Execute task AddFunc Failed, err=" + err.Error())
			}
		}
	}
	go c.Start()
	defer c.Stop()
}

func main() {
	Init()
	time.Sleep(100 * time.Second)
}
