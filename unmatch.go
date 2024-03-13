package piscine

func Unmatch(a []int) int {
	values := make(map[int]int)
	for _, char := range a {
		values[char]++
	}
	for _, char := range a {
		if values[char]%2 != 0 {
			return char
		}
	}
	return -1
}
