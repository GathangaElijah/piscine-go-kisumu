package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	maxValue := a[0]
	for _, nums := range a[1:] {
		if nums > maxValue {
			maxValue = nums
		}
	}
	return maxValue
}
