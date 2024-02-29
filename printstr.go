package piscine

import "fmt"

func PrintStr(s string) {
	aString := []string{s}
	for _, chr := range aString {
		fmt.Println(chr)
	}
}
