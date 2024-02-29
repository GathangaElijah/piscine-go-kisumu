package piscine

func UltimateDivMod(a *int, b *int) {
	if *b == 0 {
		panic("Cannot execute ")
	}

	div := *a / *b
	mod := *a % *b

	*a = div
	*b = mod
}
