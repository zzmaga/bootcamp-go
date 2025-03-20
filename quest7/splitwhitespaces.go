package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	var currentWord []rune

	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == ' ' || char == '\t' || char == '\n' {
			if len(currentWord) > 0 {
				result = append(result, string(currentWord))
				currentWord = nil
			}
		} else {
			currentWord = append(currentWord, rune(char))
		}
	}
	if len(currentWord) > 0 {
		result = append(result, string(currentWord))
	}
	return result
}
