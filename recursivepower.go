package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	if power > 1 {
		nb = nb * RecursivePower(nb, power-1)
	}
	return nb
}