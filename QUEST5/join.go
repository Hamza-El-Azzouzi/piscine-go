package piscine

func Join(strs []string, sep string) string {
	var result string
	for _, v := range strs {
		if result == "" {
			result = v
		} else {
			result = result + sep + v
		}
	}
	return result
}
