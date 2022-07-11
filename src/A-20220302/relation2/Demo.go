package main

import (
	"fmt"
)

type Relation struct {
	Uid    string `1.json:"uid"`
	Parent string `1.json:"parent"`
	Level  int    `pg:"-"  1.json:"level"`
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

	level2Relations := make(map[Relation][]Relation)
	for _, rela := range relations {
		level := rela.Level
		// level在map中是否存在
		isExists, relations := levelIsExists(level, level2Relations)
		if isExists { // 存在
			relations = append(relations, rela)

			for k, _ := range level2Relations {
				if k.Level == level && k.Uid != "" {
					delete(level2Relations, k)
				}
			}

			level2Relations[rela] = relations
		} else {
			level2Relations[rela] = []Relation{rela}
		}

	}

	fmt.Println(level2Relations)
	fmt.Println("=================")
	for k, v := range level2Relations {
		fmt.Println(k, "---", v)
	}
}

func levelIsExists(level int, level2Relations map[Relation][]Relation) (bool, []Relation) {
	for relation, relations := range level2Relations {
		l := relation.Level
		if l == level && relation.Uid != "" {
			return true, relations
		}
	}

	return false, nil
}
