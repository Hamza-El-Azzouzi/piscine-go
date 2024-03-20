package piscine

func StrRev(s string) string {
	str := []rune(s)
	len := len(str)
	for i := 0; i < len/2; i++ {
		str[i], str[len-1-i] = str[len-1-i], str[i]
	}
	return string(str)
}
