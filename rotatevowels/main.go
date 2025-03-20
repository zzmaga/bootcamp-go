package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Check if there are no arguments
	if len(os.Args) <= 1 {
		z01.PrintRune('\n')
		return
	}

	// Collect vowels from all arguments
	vowels := []rune{}
	isVowel := func(c rune) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
			c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
	}

	// Traverse arguments to extract vowels
	for _, arg := range os.Args[1:] {
		for _, char := range arg {
			if isVowel(char) {
				vowels = append(vowels, char)
			}
		}
	}

	// Reverse the vowels
	for i, j := 0, len(vowels)-1; i < j; i, j = i+1, j-1 {
		vowels[i], vowels[j] = vowels[j], vowels[i]
	}

	// Reconstruct the arguments by replacing vowels
	vowelIndex := 0
	for i, arg := range os.Args[1:] {
		for _, char := range arg {
			if isVowel(char) {
				z01.PrintRune(vowels[vowelIndex])
				vowelIndex++
			} else {
				z01.PrintRune(char)
			}
		}
		// Print a space between arguments
		if i < len(os.Args[1:])-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
