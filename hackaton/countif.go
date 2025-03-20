package piscine

func CountIf(f func(string) bool, tab []string) int {
	var result int
	for _, ok := range tab {
		if f(ok) {
			result++
		}
	}

	return result
}
