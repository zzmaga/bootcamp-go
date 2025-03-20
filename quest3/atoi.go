package piscine

func Atoi(s string) int {
	result := 0
	sign := 1
	startIndex := 0

	if len(s) > 0 && (s[0] == '+' || s[0] == '-') {
		if s[0] == '-' {
			sign = -1
		}
		startIndex = 1
	}

	for i := startIndex; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		result = result*10 + int(s[i]-'0')
	}

	return result * sign
}

// comm
