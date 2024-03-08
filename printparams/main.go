package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, value := range os.Args[1:] {
		for _, new_char := range value {
			z01.PrintRune(new_char)
		}
	}
	z01.PrintRune('\n')
}
