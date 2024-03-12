package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := len(os.Args[1:])
	if args == 0 {
		fmt.Println("File name missing")
		return
	}
	if args > 1 {
		fmt.Println("Too much arguments")
		return
	}
	filename := os.Args[1]
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Println("The filename is missing", filename)
		return
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		println("Error reading the file", err)
		return
	}
	fmt.Println(string(content))
}
