package main

import (
	"github.com/01-edu/z01"
)

func QuadA(x, y int) {
	for j := 1; j <= y; j++ {
		for i := 1; i <= x; i++ {
			if j == 1 || j == y {
				if (i == x && j == 1) || (i == x && j == y) {
					z01.PrintRune('o')
					z01.PrintRune('\n')
				} else if (i == 1 && j == 1) || (i == 1 && j == y) {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('-')
				}
			} else {
				if i == x || i == 1 {
					z01.PrintRune('|')
					if j != y && j != 1 && i == x {
						z01.PrintRune('\n')
					}
				} else {
					z01.PrintRune(' ')
				}
			}
		}
	}
}

func main() {
	QuadA(1, 5)
}
