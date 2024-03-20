package main

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	for j := 1; j <= y; j++ {
		for i := 1; i <= x; i++ {
			switch {
			case (i == 1 || i == x) && (j == 1 || j == y):
				z01.PrintRune('o')
			case i == 1 || i == x:
				z01.PrintRune('|')
			case j == 1 || j == y:
				z01.PrintRune('-')
			default:
				z01.PrintRune(' ')

			}
		}

		z01.PrintRune('\n')
	}
}

func main() {
	QuadA(10, 6)
	QuadA(5, 1)
	QuadA(1, 1)
	QuadA(1, 5)
}
