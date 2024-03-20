package main

import (
	"fmt"
)

func Split(s, sep string) []string {
	result := []string(nil)
	lastindex := 0
	lensep := len(sep)
	for i := 0; i <= len(s)-lensep; i++ {
		if s[i:i+lensep] == sep {
			result = append(result, s[lastindex:i])
			lastindex = i + lensep
			i += lensep
		} else if i == len(s)-lensep {
			result = append(result, s[lastindex:])
			lastindex = i + lensep
		}
	}
	return result
}

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}
