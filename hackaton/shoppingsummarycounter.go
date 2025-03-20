package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	counts := make(map[string]int) // Map to store word counts
	start := 0
	for i := 0; i <= len(str); i++ {
		if i == len(str) || str[i] == ' ' {
			word := str[start:i]
			counts[word]++
			start = i + 1
		}
	}
	return counts
}
