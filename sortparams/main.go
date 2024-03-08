package main

import (
	"github.com/01-edu/z01"

	"os"
)

func main() {
	params := os.Args[1:]

	for i := 0; i < len(params)-1; i++ {
		for j := 0; j < len(params)-i-1; j++ {
			if params[i] < params[i+1] {
				params[i], params[i+1] = params[i+1], params[i]
			}
		}
	}
	
	for _, chars := range params {
		for _, value := range chars {
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
	}
}
