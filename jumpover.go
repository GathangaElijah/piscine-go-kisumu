package piscine

func JumpOver(str string) string {
	if str == "" || len(str) < 3 {
		return "\n"
	}
	var newString string
	for i, value := range str {
		if (i+1)%3 == 0 {
			newString += string(value)
		}
	}
	return newString + "\n"
}
