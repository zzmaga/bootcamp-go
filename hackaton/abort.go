package piscine

func Abort(a, b, c, d, e int) int {
	sort := []int{a, b, c, d, e}
	for i := 0; i < len(sort)-1; i++ {
		for j := len(sort) - 1; j > i; j-- {
			if sort[j] > sort[j-1] {
				sort[j], sort[j-1] = sort[j-1], sort[j]
			}
		}
	}

	return sort[2]
}
