package main

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	for j := 1; j <= y; j++ {
		for i := 1; i <= x; i++ {
			if (j == 1 && i == 1 ) || (i == x && j == y) {
				z01.PrintRune('/')
			} else if (j == y && i == 1 ) || (i == x && j == 1) {
				z01.PrintRune('\\')
			} else if j == 1 || j == y {
				z01.PrintRune('*')
			} else if i == 1 || i == x {
				z01.PrintRune('*')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func main() {
	//QuadA(5, 3)
	QuadA(5, 1)
	QuadA(1, 1)
	QuadA(1, 5)
}
