package main

import "fmt"

type Boy struct {
	Name        string
	GirlFriends []string
}

func main() {
	var list []string = []string{"aa", "n", "cc", "ff"}
	for i, s := range list {
		fmt.Println(i, "------------", s)
	}

	boy := test()
	fmt.Printf("%+v\n", boy)
}

func test() Boy {
	var b Boy
	for i := 0; i < 3; i++ {
		if b.Name == "" {
			b.Name = "XM"
		}

		b.GirlFriends = append(b.GirlFriends, fmt.Sprintf("gf_%d", i))
	}
	return b
}
