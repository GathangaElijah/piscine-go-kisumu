package main

import "github.com/01-edu/z01"

func main() {
	var x rune = 48
	for i := x; i <= 57; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune(10)
}
