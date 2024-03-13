package piscine

func Abort(a, b, c, d, e int) int {
	slice := []int{a, b, c, d, e}
	lengthOfSlice := len(slice)
	for i := 0; i < lengthOfSlice-1; i++ {
		for j := 0; j < lengthOfSlice-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice[2]
}
