package piscine

func IsPrintable(s string) bool {
	for _, v := range s {
		if v < 32 || v > 126 {
			return false
		}
	}
	return true
}
