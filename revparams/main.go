package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args) - 1; i > 0; i-- {
		new_char := os.Args[i]
		for _, value := range new_char {
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
	}
}
