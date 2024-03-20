package main

import (
	"fmt"

)

func main() {
	fmt.Print(Capitalize("Hello! How are you? How+are+things+4you?"))
	fmt.Print(Capitalize("\n"))
}

func Capitalize(s string) string {
	result := []rune(s)
	isFirst := true

	for i := 0; i < len(s); i++ {
		if i == 0 {
			if result[i] >= 'a' && result[i] <= 'z'	 {
				result[i] -=32
			}
		} else {
			if (result[i-1]< 'a' || result[i-1]> 'z') && (result[i-1]< 'A' || result[i-1]> 'Z') && (result[i-1]< '0' || result[i-1]> '9') {
				isFirst = true
			} else {
				isFirst = false
			}
		}


		if isFirst {
			if result[i] >= 'a' && result[i] <= 'z'	 {
				result[i] -=32
			}
		}else{
			if result[i] >= 'A' && result[i] <= 'Z'	 {
				result[i] +=32
			}
		}

	}
	return string(result)
}