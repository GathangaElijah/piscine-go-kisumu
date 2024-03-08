package piscine

import (
	"github.com/01-edu/z01"
	"os"
)

func PrintProgramName()  {
	arguments := os.Args[0]
	for _,value := arguments[2:]{
		z01.PrintRune(value)
	}
	z01.PrintRune('\n')
}