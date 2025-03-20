package piscine

import "github.com/01-edu/z01"

func JumpOver(str string) string {
	var result string

	if len(str) < 3 {
		z01.PrintRune('\n')
	} else {
		for i := 2; i < len(str); i += 3 {
			result += string(str[i])
		}
		z01.PrintRune('\n')
	}

	return result
}
