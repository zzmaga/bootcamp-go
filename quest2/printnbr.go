package piscine

import "github.com/01-edu/z01"

func main() {
	PrintNbr(-9223372036854775808)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		for _, r := range "-9223372036854775808" {
			z01.PrintRune(r)
		}
		return
	}
	var digits []rune
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	for n > 0 {
		digit := n % 10
		digits = append(digits, rune(digit+'0'))
		n /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}
