package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	number := 1
	for i := 1; i <= power; i++ {
		number *= nb
	}
	return number
}
