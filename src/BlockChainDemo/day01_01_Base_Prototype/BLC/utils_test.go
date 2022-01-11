package BLC

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestIntToHex(t *testing.T) {
	hex := IntToHex(3)
	fmt.Println(hex)

	time := time.Now().Unix()
	fmt.Println(time)
	fmt.Printf("%x,\n%b\n", time, time)
	timeString := strconv.FormatInt(time, 2)
	fmt.Println(timeString)
	toHex := IntToHex(time)
	fmt.Println(toHex)
	fmt.Println(string(toHex))
}
