package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for a := 1; a < len(args); a++ {
		for b := a + 1; b < len(args); b++ {
			if args[a] > args[b] {
				args[a], args[b] = args[b], args[a]
			}
		}
	}
	for b := 1; b < len(args); b++ {
		for _, s := range args[b] {
			z01.PrintRune(s)
		}
		z01.PrintRune('\n')
	}
}