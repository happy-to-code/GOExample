package main

import (
	"fmt"
	"regexp"
)

func main() {
	var str string = "411402198910166320"

	reg, err := regexp.Compile("^[1-9]\\d{5}(18|19|20|(3\\d))\\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\\d{3}[0-9Xx]$")
	if err != nil {
		panic(err)
	}

	if reg.MatchString(str) {
		fmt.Println("匹配成功")

		// submatch := reg.FindStringSubmatch(str)
		// fmt.Println(submatch)
		// fmt.Println("date is ", submatch[2])
	}

}
