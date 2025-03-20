package piscine

// SortWordArr sorts a slice of strings by their ASCII values in ascending order
func SortWordArr(a []string) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Compare two adjacent elements
			if a[j] > a[j+1] {
				// Swap the elements if they are in the wrong order
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
