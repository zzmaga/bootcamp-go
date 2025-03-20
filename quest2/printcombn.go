package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	combination := make([]int, n)
	for i := 0; i < n; i++ {
		combination[i] = i
	}

	max := 10 - n

	for {
		for _, digit := range combination {
			z01.PrintRune(rune(digit + '0'))
		}

		if combination[0] == max && combination[n-1] == 9 {
			break
		}

		z01.PrintRune(',')
		z01.PrintRune(' ')

		for i := n - 1; i >= 0; i-- {
			if combination[i] < 9-(n-1-i) {
				combination[i]++
				for j := i + 1; j < n; j++ {
					combination[j] = combination[j-1] + 1
				}
				break
			}
		}
	}

	z01.PrintRune('\n')
}
