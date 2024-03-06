package piscine

func ToUpper(s string) string {
	var new_string string
	for _, value := range s {
		if value >= 'a' && value <= 'z' {
			new_string += string(value - 32)
		} else {
			new_string += string(value)
		}
	}
	return new_string
}
