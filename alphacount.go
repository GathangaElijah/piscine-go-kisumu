package piscine

func AlphaCount(s string) int {
	count := 0
	for _, values := range s {
		if values >= 'a' && values <= 'z' || values >= 'A' && values <= 'Z' {
			count++
		}
	}
	return count
}
