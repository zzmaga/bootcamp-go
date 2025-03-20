package piscine

func LastRune(s string) rune {
	var l rune
	for _, r := range s {
		for i := len(s); i >= 0; i-- {
			l = r
		}
	}
	return l
}

/* func LastRune(s string) rune {
 	if len(s) == 0 {
 		return 0
 	}
 	return rune(s[len(s)-1])
}
*/
