package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[0]
	for _, value := range arguments[2:] {
		z01.PrintRune(value)
	}
	z01.PrintRune('\n')
}
