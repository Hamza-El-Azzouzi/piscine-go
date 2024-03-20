package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintStr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func PrintNbr(n int) {
	nb := '0'
	subnb := '0'
	for i := 1; i <= n/10; i++ {
		nb++
	}
	z01.PrintRune(nb)
	for i := 1; i <= n%10; i++ {
		subnb++
	}
	z01.PrintRune(subnb)
}

func main() {
	points := &point{}
	setPoint(points)
	PrintStr("x = ")
	PrintNbr(points.x)
	PrintStr(", y = ")
	PrintNbr(points.y)
	PrintStr("\n")
}
