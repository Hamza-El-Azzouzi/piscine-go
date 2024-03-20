package piscine

func Any(f func(string) bool, a []string) bool {
	var result bool
	for _, v := range a {
		if f(v) == true {
			result = f(v)
			break
		} else {
			result = f(v)
		}
	}
	return result
}
