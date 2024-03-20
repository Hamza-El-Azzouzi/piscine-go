package piscine

func BasicAtoi2(s string) int {
	result := 0
	for _, nb := range s {
		if nb < '0' || nb > '9' {
			return 0
		}
		result = result*10 + int(nb-'0')
	}
	return result
}
