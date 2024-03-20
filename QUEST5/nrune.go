package piscine

func NRune(s string, n int) rune {
	if n-1 < 0 || len(s) < n {
		return 0
	} else {
		result := []rune(s)
		return result[n-1]
	}
}
