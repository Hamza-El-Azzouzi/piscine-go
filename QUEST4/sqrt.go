package piscine

func Sqrt(nb int) int {
	var result int
	for i := 1; i <= nb; i++ {
		if nb== i*i {
			result = i
		}
	}
	return result
}
