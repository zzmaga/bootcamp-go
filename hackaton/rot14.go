package piscine

func Rot14(s string) string {
	var result []rune

	for _, r := range s {

		if r >= 'a' && r <= 'z' {
			if r < 'n' {
				r = r + 14
			} else {
				r = r - 12
			}
		}
		if r >= 'A' && r <= 'Z' {
			if r < 'N' {
				r = r + 14
			} else {
				r = r - 12
			}
		}
		result = append(result, r)
	}

	return string(result)
}
