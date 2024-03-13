package piscine

func Compact(ptr *[]string) int {
	strValues := *ptr
	count := 0
	lenOfString := len(strValues)
	for i := 0; i < lenOfString; i++ {
		if strValues[i] != "" {
			strValues[count] = strValues[i]
			count++
		}
	}
	*ptr = strValues[:count]
	return count
}
