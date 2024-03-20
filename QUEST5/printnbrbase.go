package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
}
func PrintNbrBase(n int, base string) {
	if ValidateBase(base) {
		length := len(base)
		sign := 1
		rbase := []rune(base)
		if n < 0 {
			z01.PrintRune('-')
			sign = -1
		}
		if n < length && n >= 0 {
			z01.PrintRune(rbase[n])
		} else {
			PrintNbrBase(sign*(n/length), base)
			z01.PrintRune(rbase[sign*(n%length)])
		}
	} else {
		z01.PrintRune('N')
		z01.PrintRune('V')
	}
}
func ValidateBase(base string) bool {
	Len := len(base)
	if Len <= 1 {
		return false
	} else if string(base[0]) == "-" || string(base[0]) == "+" {
		return false
	} else {
		for i := 0; i < Len; i++ {
			for j := 0; j < Len; j++ {
				if (base[i] == base[j]) && !(j == i) {
					return false
				}
			}
		}
		return true
	}
}