package piscine

func AppendRange(min, max int) []int {
	params := []int{}
	for i := min; i < max; i++ {
		params = append(params, i)
	}
	return params
}
