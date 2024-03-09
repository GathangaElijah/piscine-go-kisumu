package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	params := make([]int, max-min)
	for i := min; i < max; i++ {
		params[i] = min + i
	}
	return params
}
