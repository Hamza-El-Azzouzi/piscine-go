package piscine

func Map(f func(int) bool, a []int) []bool {
	var result []bool
	for _,v := range a {
		result = append(result, f(v))
	}
	return result
}

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(IsPrime, a)
	fmt.Println(result)
}