package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascending := true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			ascending = false
			break
		}
	}

	descending := true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			descending = false
			break
		}
	}

	return ascending || descending
}
