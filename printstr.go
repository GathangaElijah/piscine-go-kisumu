package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, chr := range s {
		z01.PrintRune(chr)
	}
	z01.PrintRune(10)
}
