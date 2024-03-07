package piscine

func TrimAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	var new_int int
	var sign int = 1
	start_looking := false

	for _, value := range s {
		if value >= '0' && value <= '9' {
			start_looking = true
			new_int = new_int*10 + int(value-'0')
		} else if value == '-' && !start_looking {
			sign = -1
		} else if start_looking && (value < '0' || value > '9') {
			continue
		}
	}
	return new_int * sign
}
