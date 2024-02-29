package piscine

func UltimateDivMod(a *int, b *int) {
	if b == 0 {
		panic("Cannot execute b is 0")
	}

	*a = *a / *b
	*b = *a % *b
}
