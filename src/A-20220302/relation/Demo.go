package main

import "fmt"

type Relation struct {
	Uid    string `json:"uid"`
	Parent string `json:"parent"`
	Level  int    `pg:"-"  json:"level"`
}

func main() {
	relations := []Relation{
		{
			Uid:    "7",
			Parent: "8",
			Level:  0,
		},
		{
			Uid:    "3",
			Parent: "7",
			Level:  1,
		},
		{
			Uid:    "2",
			Parent: "4",
			Level:  2,
		},
		{
			Uid:    "5",
			Parent: "4",
			Level:  2,
		},
		{
			Uid:    "1",
			Parent: "4",
			Level:  2,
		},
		{
			Uid:    "6",
			Parent: "5",
			Level:  3,
		},
		{
			Uid:    "4",
			Parent: "3",
			Level:  3,
		},
	}
	fmt.Println(relations)

	level2Ids := make(map[int][]string)
	for _, rela := range relations {
		level := rela.Level
		uid := rela.Uid
		ids, _ := level2Ids[level]
		ids = append(ids, uid)
		level2Ids[level] = ids
	}

	fmt.Println(level2Ids)
	fmt.Println("=================")
	for k, v := range level2Ids {
		fmt.Println(k, "---", v)
	}
}
