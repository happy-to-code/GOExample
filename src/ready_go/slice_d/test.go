package main

import "fmt"

func main1() {
	s := []int{5}
	fmt.Println("cap1:", cap(s))
	s = append(s, 7)
	fmt.Println("cap2:", cap(s))
	s = append(s, 9)
	fmt.Println("cap3:", cap(s))

	x := append(s, 11)
	fmt.Println("cap4:", cap(s))
	y := append(s, 12)
	fmt.Println("cap5:", cap(s))
	fmt.Printf("s:%v, x:%v, y:%v", s, x, y)
	// s:[5 7 9], x:[5 7 9 12], y:[5 7 9 12]
}

func main() {
	s := []int{5, 7, 9}
	fmt.Println("cap1:", cap(s))
	x := append(s, 11)
	fmt.Println("cap2:", cap(s))
	y := append(s, 12)
	fmt.Println("cap3:", cap(s))
	fmt.Printf("s:%v, x:%v, y:%v", s, x, y)
	// s:[5 7 9], x:[5 7 9 11], y:[5 7 9 12]
}
