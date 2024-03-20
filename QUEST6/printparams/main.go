package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		params := []rune(os.Args[i])
		for j := 0; j < len(params); j++ {
			z01.PrintRune(params[j])
		}
		z01.PrintRune('\n')
	}
}
