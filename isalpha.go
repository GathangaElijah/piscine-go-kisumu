package piscine

func IsAlpha(s string) bool {
	for _, value := range s {
		if !((value >= 'A' && value <= 'Z') ||
			(value >= 'a' && value <= 'z') ||
			(value >= '0' && value <= '9')) {
			return false
		}
	}
	return true
}
