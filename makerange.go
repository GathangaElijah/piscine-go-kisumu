package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	params := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		params[i] = min + i
	}
	return params
}
