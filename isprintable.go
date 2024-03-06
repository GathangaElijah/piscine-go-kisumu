package piscine

func IsPrintable(s string) bool {
	printchars := []rune(s)
	for _, value := range printchars {
		if value >= 0 && value <= 31 {
			return false
		}
	}
	return true
}
