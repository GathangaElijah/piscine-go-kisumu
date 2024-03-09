package piscine

func MakeRange(min, max int) []int {
	var nil_params []int
	if min >= max {
		return nil_params
	}
	params := make([]int, 0)
	for i := min; i < max; i++ {
		params = make([]int, i)
	}
	return params
}
