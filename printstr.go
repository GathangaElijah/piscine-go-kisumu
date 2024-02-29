package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	aString := []string{s}
	for _, chr := range aString {
		z01.PrintRune(chr)
	}
}
