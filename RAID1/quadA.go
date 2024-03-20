package main

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		// the first loop is for axes j (3amod j li fih y)
		for j := 1; j <= y; j++ {
			// the second loop is for axes i(3amod i li fih x)
			for i := 1; i <= x; i++ {
				// this is the first condition to select the first line and the last line(y) in j axes
				if j == 1 || j == y {
					// this is the second if for i axes  to print 'o' (l9nat b 4)
					if i == 1 || i == x {
						z01.PrintRune('o')
						// else to print '-' between 'o' in the first and the last line
					} else {
						z01.PrintRune('-')
					}
					// this else for the first condition to select between the first and the last line in j axes
				} else {
					// this if to print pipes in the first(1) and the last line(i) in i axes
					if i == 1 || i == x {
						z01.PrintRune('|')
						// else to print spaces between pipes
					} else {
						z01.PrintRune(' ')
					}
				}
			}
			// to create a new line everytime you finish from a line
			z01.PrintRune('\n')
		}
	}
}

func main() {
	QuadA(0, 3)
	//QuadA(5, 1)
	//QuadA(1, 1)
	//QuadA(1, 5)
}
