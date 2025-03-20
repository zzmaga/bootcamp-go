package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	for x := 0; x*x <= nb; x++ {
		if x*x == nb {
			return x
		}
	}
	return 0
}
