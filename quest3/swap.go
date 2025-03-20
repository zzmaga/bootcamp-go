package piscine

func Swap(a *int, b *int) {
	k := *b
	*b = *a
	*a = k
}
