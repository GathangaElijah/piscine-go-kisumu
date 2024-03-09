package piscine

func ConcatParams(args []string) string {
	var params string
	for i, value := range args {
		if i > 0 {
			params += "\n"
		}
		params += value
	}
	return params
}
