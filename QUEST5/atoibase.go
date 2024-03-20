package main
import (
	"fmt"
)
func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

func Index(s string, toFind string) int {
	for i := 0; i <= (len(s) - len(toFind)); i++ {
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}
	return -1
}
func ValidateBase(base string) bool {
	if len(base) < 2 || Index(base, "+") >= 0 || Index(base, "-") >= 0 {
		return false
	} else {
		for i, char1 := range base {
			for j, char2 := range base {
				if i != j && char1 == char2 {
					return false
				}
			}
		}
		return true
	}
}
func IterativePower(nb int, power int) int {
	result := 1
	if power < 0 {
		return 0
	}
	for i := 1; i <= power; i++ {
		result = result * nb
	}
	return result
}
func AtoiBase(s string, base string) int {
	if !ValidateBase(base) {
		return 0
	}
	result := 0
	for i, ch := range s {
		result += Index(base, string(ch)) * IterativePower(len(base), len(s)-1-i)
	}
	return result
}
