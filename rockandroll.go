package piscine

func RockAndRoll(n int) string {
	if n < 0 {
		error := "error: number is negative\n"
		return error
	} else if n%2 == 0 && n%3 == 0 {
		message3 := "rock and roll\n"
		return message3
	} else if n%2 == 0 {
		message1 := "rock\n"
		return message1
	} else if n%3 == 0 {
		message2 := "rock\n"
		return message2
	} else if n%2 == 0 && n%3 == 0 {
		message3 := "rock and roll\n"
		return message3
	} else {
		return "non divisible\n "
	}
}