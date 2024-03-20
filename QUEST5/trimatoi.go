package piscine

func TrimAtoi(s string) int {
	result := 0
	t := 1
	for _, v := range s {
		if v >= '0' && v <= '9' {
			t = 1
			break
		} else if v == '-' {
			t = -1
			break
		}
	}
	for _, v := range s {
		if v >= '0' && v <= '9' {
			result = result*10 + int(v-48)
		}
	}
	return result * t
}
