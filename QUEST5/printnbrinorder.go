package piscine

import "github.com/01-edu/z01"

func Sortnbr(nb []int) {
	for i := range nb {
		for j := range nb {
			if nb[i] < nb[j] {
				nb[i], nb[j] = nb[j], nb[i]
			}
		}
	}
}

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	nb := []int{}
	for n > 0 {
		nb = append(nb, n%10)
		n = n / 10
	}
	Sortnbr(nb)
	for i := 0; i < len(nb); i++ {
		z01.PrintRune(rune(nb[i] + '0'))
	}
}
