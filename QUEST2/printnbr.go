package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}
func PrintNbr(n int) {
	nbr := []int{}
	if n < 0 {
		z01.PrintRune('-')
		n = -1 * n
	} else if n == 0 {
		z01.PrintRune('0')
	}
	for n > 0 {
		nbr = append(nbr, n%10)
		n = n / 10
	}
	for i := len(nbr) - 1; i >= 0; i-- {
		z01.PrintRune(rune(48 + nbr[i]))
	}
}
