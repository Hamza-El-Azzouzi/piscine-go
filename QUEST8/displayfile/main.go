package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}
	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	content, _ := ioutil.ReadFile(args[0])
	fmt.Print(string(content))
}
