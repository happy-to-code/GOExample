package main

import "fmt"

func main() {
	var a int64
	fmt.Println(a, "----")
	m := make(map[string]Member, 2)
	m["account1"] = Member{"name1", "account1"}
	m["account2"] = Member{"name2", "account2"}

	fmt.Println(m)
	fmt.Println("------------------")
	member, ok := m["account1"]
	fmt.Println(member, "--", ok)
	fmt.Println("-->", member.Name)

}

type Member struct {
	Name    string
	Account string
}
