package main

import (
	"fmt"
)
func Prime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 || n == 3{
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5
	for i*i<=n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i+=6
	}
	return true

}
func FindNextPrime(nb int) int {
	if Prime(nb) {
		return nb
	}
	for i := 1 ; ; i++ {
		if Prime(nb+i) {
			return nb+i
		}
	}
}


func main() {
	fmt.Println(FindNextPrime(5))
	fmt.Println(FindNextPrime(4))
}