package main

import "log"

func main() {
	var i = 0
	testReturn(i)

	log.Println("end ……")
}

func testReturn(a int) {
	if a == 0 {
		log.Println("a is zero")
		return
	}
	log.Println("a + 1 = ", a+1)
}
