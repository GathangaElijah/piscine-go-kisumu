package main

import (
	"fmt"
	"os"
)

func main() {
	listOfArgs := os.Args[1:]
	for _, arg := range listOfArgs {
		if arg == "01" || arg == "galaxy" || arg == "galaxy 01" {
			fmt.Println("Alert!!!")
		} else {
			println()
		}
	}
}
