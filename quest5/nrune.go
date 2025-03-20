package piscine

func NRune(s string, n int) rune {
	if n < 1 || n > len(s) {
		return 0
	}
	var zhana []rune
	for _, r := range s {
		zhana = append(zhana, r)
	}
	return zhana[n-1]
}
