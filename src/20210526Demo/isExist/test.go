package main

import "fmt"

func main() {
	ledgerIdsFromDB := []string{"aa", "bb", "cc", "dd"}

	temp := ""
	idIsExist := ledgerIdIsExist1(ledgerIdsFromDB, temp)
	fmt.Println("sss:", idIsExist)
}
func ledgerIdIsExist1(ledgerIdsFromDB []string, ledgerId string) bool {
	for _, idFromDB := range ledgerIdsFromDB {
		if ledgerId == idFromDB {
			return true
		}
	}
	return false
}
