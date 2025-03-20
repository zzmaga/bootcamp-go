package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, ok := range a {
		result[i] = f(ok)
	}

	return result
}
