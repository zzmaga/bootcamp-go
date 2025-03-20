package piscine

func BasicAtoi2(s string) int {
	result := 0

	for _, r := range s {
		if r >= '0' && r <= '9' {
			result = result*10 + int(r-'0')
		} else {
			return 0
		}
	}

	return result
}

// comm
