package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	runes := []rune(name)
	for i := 2; i < len(runes); i++ {
		z01.PrintRune(runes[i])
	}
	z01.PrintRune('\n')
}
