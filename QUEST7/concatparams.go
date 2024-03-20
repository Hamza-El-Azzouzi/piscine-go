package piscine

func ConcatParams(args []string) string {
	r := ""
	for i, v := range args {
		r = r + v
		if i < len(args)-1 {
			r += "\n"
		}
	}
	return r
}
