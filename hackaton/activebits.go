package piscine

func ActiveBits(n int) int {
	var bek int = 0
	for n > 0 {
		bek += n & 1
		n >>= 1
	}
	return bek
}
