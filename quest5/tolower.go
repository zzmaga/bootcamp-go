package piscine

func ToLower(s string) string {
	result := []rune(s)
	for i := 0; i < len(result); i++ {
		if 'A' <= result[i] && result[i] <= 'Z' {
			result[i] = result[i] + ('a' - 'A')
		}
	}
	return string(result)
}
