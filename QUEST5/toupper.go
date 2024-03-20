package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			runes[i] = runes[i] - 32
		}
	}
	return string(runes)
}
