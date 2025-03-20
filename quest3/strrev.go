package piscine

func StrRev(s string) string {
	var sozder []rune
	for i := len(s) - 1; i >= 0; i-- {
		sozder = append(sozder, rune(s[i]))
	}
	return string(sozder)
}

// comm
