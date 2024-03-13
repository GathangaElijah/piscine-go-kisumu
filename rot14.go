package piscine

func rot14(s string) string {
	var newString string
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			newString += string((char+14)%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			newString += string((char+14)%26 + 'Z')
		} else {
			newString += string(char)
		}
	}
	return newString
}
