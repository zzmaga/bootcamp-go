package piscine

func TrimAtoi(s string) int {
	var result int
	var sign int = 1
	var foundNumber bool = false

	for _, char := range s {
		if char == '-' && !foundNumber {
			sign = -1
		} else if char >= '0' && char <= '9' {
			foundNumber = true
			result = result*10 + int(char-'0')
		}
	}

	return result * sign
}
