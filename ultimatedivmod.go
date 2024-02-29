package piscine

func UltimateDivMod(a *int, b *int) {
	if *b == 0 {
		panic("Cannot execute ")
	}

	*a = *a / *b
	*b = *a % *b
}
