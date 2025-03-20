package piscine

func Any(f func(string) bool, a []string) bool {
	var result bool
	for _, ok := range a {
		if f(ok) {
			return true
		}
	}

	return result
}
