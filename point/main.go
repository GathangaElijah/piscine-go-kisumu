package main

import (
	"github.com/01-edu/z01"
)

var stringOutput = "x = 42, y = 21"

func main() {
	for _, char := range stringOutput {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
