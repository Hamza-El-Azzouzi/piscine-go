package piscine

func SplitWhiteSpaces(s string) []string {
	var slice []string
	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '\n' && s[i] != '\t' {
			word += string(s[i])
		} else if word != "" {
			slice = append(slice, word)
			word = ""
		}
	}
	if word != "" {
		slice = append(slice, word)
	}
	return slice
}
