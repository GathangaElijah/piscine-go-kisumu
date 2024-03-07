package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	no_of_digits := make([]int, 0, 10)
	for i := range no_of_digits {
		for j := range no_of_digits {
			if no_of_digits[i] < no_of_digits[j] {
				no_of_digits[i], no_of_digits[j] = no_of_digits[j], no_of_digits[i]
			}
		}
	}
	for _, values := range no_of_digits {
		z01.PrintRune(rune('0' + values))
	}
}
