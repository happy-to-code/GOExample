package main

import "fmt"

type Monster struct {
	Name string
}

// NewMonster 我们来到一个黑暗的世界，这个世界中有一个邪恶的怪兽
func NewMonster() Monster {
	return Monster{Name: "kitty"}
}

type Player struct {
	Name string
}

func NewPlayer(name string) Player {
	return Player{Name: name}
}

type Mission struct {
	Player  Player
	Monster Monster
}

func NewMission(p Player, m Monster) Mission {
	return Mission{p, m}
}

func (m Mission) Start() {
	fmt.Printf("%s defeats %s, world peace!\n", m.Player.Name, m.Monster.Name)
}

func main() {
	monster := NewMonster()
	player := NewPlayer("dj")
	mission := NewMission(player, monster)

	mission.Start()
}
