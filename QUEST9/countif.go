package piscine

func CountIf(f func(string) bool, tab []string) int {
	var result []bool
	for _, v := range tab {
		if f(v) == true {
			result = append(result, f(v))
		}
	}
	return len(result)
}
