package piscine

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if nb == 0 {
		return 0
	}
	if power < 0 || nb > 20 {
		return 0
	}
	result := 1
	eky := nb
	for i := 1; i <= power; i++ {
		result = result * eky
	}
	return result
}
