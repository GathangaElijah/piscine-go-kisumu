package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	upper := false
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	if len(args) == 0 {
		return
	}
	for _, arg := range args {
		n, ok := Atoi(arg)
		if !ok || n < 1 || n > 26 {
			z01.PrintRune(' ')
			continue
		}

		letter := rune('a' + n - 1)
		if upper {
			letter = rune('A' + n - 1)
		}
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}

func Atoi(s string) (int, bool) {
	var result int
	var negative bool

	if s[0] == '-' {
		negative = true
		s = s[1:]
	}

	for _, digit := range s {
		if digit < '0' || digit > '9' {
			return 0, false
		}
		result = result*10 + int(digit-'0')
	}

	if negative {
		result = -result
	}
	return result, true
}
