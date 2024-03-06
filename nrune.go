package piscine

func NRune(s string, n int) rune {
	allrunes := []rune(s)
	if n >= 1 && n <= len(allrunes) {
		return allrunes[n-1]
	}
	return 0
}
