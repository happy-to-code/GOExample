package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(checkUsername("kdlsdf234"))
}

func checkUsername(username string) bool {
	const pattern = `^[a-z0-9_-]{3,16}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(username)
}
