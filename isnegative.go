package main

import (
	"fmt"
)

func IsNegative(nb int) {
	if nb < 0 {
		fmt.Println("T")
	} else {
		fmt.Println("F")
	}
}

func main() {
	IsNegative(-4)
}
