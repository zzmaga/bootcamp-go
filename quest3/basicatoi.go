package piscine

func BasicAtoi(s string) int {
	result := 0
	for _, r := range s {
		if r >= '0' && r <= '9' {
			san := int(r - '0')
			result = result*10 + san
		}
	}
	return result
}

// comm
