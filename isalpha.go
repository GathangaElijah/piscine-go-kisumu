package piscine

func IsAlpha(s string) bool {
	for _, value := range s {
		// check out characters is uppercase
		if !((value >= 'A' && value <= 'Z') ||
			//checks out characters is lowercase
			(value >= 'a' && value <= 'z') ||
			// checks if the character is  a numeric
			(value >= '0' && value <= '9')) {
			return false
		}
	}
	return true
}
