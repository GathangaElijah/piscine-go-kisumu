package piscine

func StrRev(s string) string {
	mystring := []rune(s)
	length := len(mystring)
	reversed := make([]rune, length)
	for i, str := range mystring {
		reversed[length-i-1] = str
	}
	return string(reversed)
}
