package main

import (
	"fmt"
	"github.com/go-pg/pg/v10"
)

type Job struct {
	tableName struct{} `pg:"cron.job"`
	Jobid     int64    `json:"jobid"`
	Schedule  string   `json:"schedule"`
	Command   string   `json:"command"`
	Nodename  string   `json:"nodename"`
	Nodeport  int      `json:"nodeport"`
	Database  string   `json:"database"`
	Username  string   `json:"username"`
	Active    string   `json:"active"`
	Jobname   string   `json:"jobname"`
}

func main() {
	var job = Job{
		Schedule: "0 30 * * *",
		Command:  "test",
		Nodename: "localhost",
		Nodeport: 5432,
		Database: "--postgres--",
		Username: "postgres",
		Active:   "t",
		Jobname:  "test_job_name7",
	}

	db := pg.Connect(&pg.Options{
		Database: "postgres",
		Addr:     fmt.Sprintf("%s:%d", "10.1.13.103", 5432),
		User:     "web_user",
		Password: "web@2022",
	})

	fmt.Println(job)
	// result, err := db.Model(&job).Insert()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("result:", result)

	var s = "select schedule_in_database(\"test007\", \"0 30 * * *\", \"test\", \"postgres\", \"postgres\", \"t\")"
	// result, err := db.Exec("schedule_in_database", "test007", "0 30 * * *", "test", "postgres---", "postgres", "t")
	fmt.Println("===>", s)
	result, err := db.Exec(s)
	if err != nil {
		panic(err)
	}
	fmt.Println("-->", result)

}
