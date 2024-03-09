package piscine

func Split(s, sep string) []string {
	new_string := make([]string, 0)
	current_string := ""
	sep_len := len(sep)
	for i := 0; i < len(s); i++ {
		if i+sep_len <= len(s) && s[i:i+len(sep)] == sep {
			new_string = append(new_string, current_string)
			current_string = ""
			i += len(sep) - 1
		} else {
			current_string += string(s[i])
		}
	}
	if current_string != "" {
		new_string = append(new_string, current_string)
	}
	return new_string
}
