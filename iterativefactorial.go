package piscine

func IterativeFactorial(nb int) int {
	if nb <= 0 {
		return 0
	}

	factorial := 1

	for i := 1; i <= nb; i++ {
		factorial *= i
	}
	return factorial
}
