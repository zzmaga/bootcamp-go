package piscine

func StringToIntSlice(str string) []int {
	var bek []int
	for _, r := range str {
		bek = append(bek, int(r))
	}

	return bek
}
