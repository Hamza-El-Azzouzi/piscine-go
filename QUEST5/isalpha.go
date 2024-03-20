package piscine

func IsAlpha(s string) bool {
	for _, v := range s {
		if (v < 'A' || v > 'Z') && (v < 'a' || v > 'z') && (v < '0' || v > '9') {
			return false
		}
	}
	return true
}
