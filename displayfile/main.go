package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := len(os.Args[1:])
	if args == 0 {
		fmt.Printf("File name missing\n")
		return
	}
	if args > 1 {
		fmt.Printf("Too many arguments\n")
		return
	}
	filename := os.Args[1]
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Printf("The filename is missing%v\n", filename)
		return
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading the file%v\n", err)
		return
	}
	fmt.Printf(string(content))
}
