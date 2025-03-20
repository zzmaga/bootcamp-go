package piscine

func SortIntegerTable(table []int) {
	n := len(table)
	for i := 0; i < n-1; i++ {
		for j := n - 1; j > i; j-- {
			if table[j] < table[j-1] {
				table[j], table[j-1] = table[j-1], table[j]
			}
		}
	}
}
