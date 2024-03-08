package piscine

func Capitalize(s string) string {
	var output []rune
	makeCapital := true
	for _, value := range s {
		if makeCapital && isLetterOrDigit(value) {
			value = toUpperCase(value)
			makeCapital = false
		} else if !isLetterOrDigit(value) {
			makeCapital = true
		} else if !makeCapital && isLetter(value) {
			value = toLowerCase(value)
		}
		output = append(output, value)
	}
	return string(output)
}

func isLetterOrDigit(value rune) bool {
	return (value >= 'a' && value <= 'z') || (value >= 'A' && value <= 'Z') || (value >= '0' && value <= '9')
}

func isLetter(value rune) bool {
	return (value >= 'a' && value <= 'z') || (value >= 'A' && value <= 'Z')
}

func toUpperCase(value rune) rune {
	if value >= 'a' && value <= 'z' {
		return value - 32
	}
	return value
}

func toLowerCase(value rune) rune {
	if value >= 'A' && value <= 'Z' {
		return value + 32
	}
	return value
}