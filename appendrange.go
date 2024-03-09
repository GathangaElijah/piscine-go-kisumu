package piscine

func AppendRange(min, max int) []int {
	params := []int{}
	nil_params := []int(nil)
	for i := min; i < max; i++ {
		if min >= max {
			return nil_params
		} else {
			params = append(params, i)
		}
	}
	return params
}
