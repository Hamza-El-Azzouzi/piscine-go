package piscine

func BasicAtoi(s string) int {
	r := 0
	for _, v := range s {
		r = r*10 + int(v-'0')
	}
	return r
}
