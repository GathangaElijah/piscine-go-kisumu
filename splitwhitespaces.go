package piscine

func SplitWhiteSpaces(s string) []string {
	var params_string []string
	var current_string string
	for _, value := range s {
		if value == '\t' || value == ' ' || value == '\r' || value == '\n' {
			if current_string != "" {
				params_string = append(params_string, current_string)
				current_string = ""
			}
		} else {
			current_string += string(value)
		}
	}
	if current_string != "" {
		params_string = append(params_string, current_string)
	}
	return params_string
}
