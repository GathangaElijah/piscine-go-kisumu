package piscine

func IsUpper(s string) bool {
	for _, values := range s {
		if values < 'A' && values < 'Z' {
			return false
		}
	}
	return true
}
