package piscine

func StringToArray(str string) []string {
	var result []string
	var word string

	for _, r := range str {
		if r == ' ' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(r)
		}
	}
	if word != "" {
		result = append(result, word)
	}

	return result
}
