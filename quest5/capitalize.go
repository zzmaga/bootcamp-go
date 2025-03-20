package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	inWord := false

	for i, char := range runes {
		if isAlphaNumeric(char) {
			if !inWord {
				runes[i] = toUpper(char)
				inWord = true
			} else {
				runes[i] = toLower(char)
			}
		} else {
			inWord = false
		}
	}
	return string(runes)
}

func isAlphaNumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}

func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - ('a' - 'A')
	}
	return r
}
