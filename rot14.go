package piscine

func Rot14(s string) string {
	var newString string
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			newString += string(((char-'a')+14)%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			newString += string(((char-'A')+14)%26 + 'A')
		} else {
			newString += string(char)
		}
	}
	return newString
}
