package piscine

func MakeRange(min, max int) []int {
	params := make([]int, 0)
	for i := min; i < max; i++ {
		if min >= max {
			return params
		} else {
			params = make([]int, i)
		}
	}
	return params
}
