package main

import (
	"A-20220302/cobra2/cmd"
	"fmt"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
