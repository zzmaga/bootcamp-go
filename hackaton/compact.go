package piscine

func Compact(ptr *[]string) int {
	if ptr == nil {
		return 0
	}

	var str []string
	num := 0
	for _, d := range *ptr {
		if d != "" {
			str = append(str, d)
			num++
		}
	}
	*ptr = str
	return num
}
