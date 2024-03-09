package piscine

func AppendRange(min, max int) []int {
	params := []int{}
	for i := min; i < max; i++ {
		if min >= max {
			return params
		} else {
			params = append(params, i)
		}
	}
	return params
}
