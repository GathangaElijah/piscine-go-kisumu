package piscine

func MakeRange(min, max int) []int {
	params := make([]int, 0)
	var nil_params []int
	for i := min; i < max; i++ {
		if min >= max {
			return nil_params
		} else {
			params = make([]int, i)
		}
	}
	return params
}
