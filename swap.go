package piscine

func Swap(a *int, b *int) {
	//*a = *b
	//*b = *a
	
	a = *&a
	b = *&b

	a = b
	b = a
}
