package piscine

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	newString := strs[0]
	for _, value := range strs[1:] {
		newString += sep + value
	}
	return newString
}
