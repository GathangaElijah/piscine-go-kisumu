package piscine

func StringToIntSlice(str string) []int {
	var newStringInt []int
	for _, value := range str {
		newStringInt = append(newStringInt, int(value))
	}
	return newStringInt
}
