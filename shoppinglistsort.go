package piscine

func ShoppingListSort(slice []string) []string {
	lengthOfString := len(slice)
	for i := 0; i < lengthOfString-1; i++ {
		for j := 0; i < lengthOfString; j++ {
			if len(slice[i]) > len(slice[j]) {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}
