package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, word := range a {
		for _, r := range word {
			_ = z01.PrintRune(r)
		}

		_ = z01.PrintRune('\n')
	}
}
