package piscine

func DescendAppendRange(max, min int) []int {
	var newRange []int
	if max <= min {
		return []int{}
	}
	for i := max; i > min; i-- {
		newRange = append(newRange, i)
	}
	return newRange
}
