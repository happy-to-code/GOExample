package main

import "fmt"

func main() {
	retry()
}

func retry() {
	err := print()
	if err != nil {
		for i := 0; i < 3; i++ {
			err2 := print()
			if i == 2 {
				err2 = nil
			}
			if err2 == nil {
				break
			}
		}
	}

}

func print() error {
	return fmt.Errorf("%s", "sdfsdfsdfsdfsd")
}
