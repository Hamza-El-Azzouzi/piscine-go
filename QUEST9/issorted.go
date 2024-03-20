package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) > 0 {
			for i := 2; i < len(a); i++ {
				if f(a[i-1], a[i]) < 0 {
					return false
				}
			}
		} else if f(a[i-1], a[i]) > 0 {
			for i := 2; i < len(a); i++ {
				if f(a[i-1], a[i]) > 0 {
					return false
				}
			}
		}
	}
	return true
}
