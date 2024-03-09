package piscine

func MakeRange(min, max int) []int {
	var nil_params []int
	if min >= max {
		return nil_params
	}
	params := make([]int, max-min+1)
	for i := min; i < max-min; i++ {
		params[i] = min + i
	}
	return params
}
