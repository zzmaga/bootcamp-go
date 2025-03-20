package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	decimal := 0
	for i := 0; i < len(nbr); i++ {
		val := -1
		for j := 0; j < len(baseFrom); j++ {
			if nbr[i] == baseFrom[j] {
				val = j
				break
			}
		}
		if val == -1 {
			return ""
		}
		decimal = decimal*len(baseFrom) + val
	}

	if decimal == 0 {
		return string(baseTo[0])
	}

	result := ""
	baseLen := len(baseTo)
	for decimal > 0 {
		result = string(baseTo[decimal%baseLen]) + result
		decimal /= baseLen
	}

	return result
}
