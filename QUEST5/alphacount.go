package piscine

func AlphaCount(s string) int {
	count := 0
	for _, v := range s {
		if v >= 'A' && v <= 'Z' || v >= 'a' && v <= 'z' {
			count++
		}
	}
	return count
}
