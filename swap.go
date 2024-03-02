package piscine

func Swap(a *int, b *int) {
	var value1 int
	var value2 int

	value1 = *a
	value2 = *b

	value1 = value2
	value2 = value1
}
