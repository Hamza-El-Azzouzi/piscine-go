package piscine

func Atoi(s string) int {
	r := 0
	t := 1
	for i, nb := range s {
		if nb == '-' && i == 0 {
			t = -1
			continue
		}
		if nb == '+' && i == 0 {
			t = 1
			continue
		}
		if nb < '0' || nb > '9' {
			return 0
		}
		r = r*10 + int(nb-'0')
	}
	return r * t
}
