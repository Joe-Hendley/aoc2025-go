package search

func Binary(lower, upper int, f func(i int) bool, direction func(i int) bool) int {
	mid := (lower + upper) / 2
	if f(mid) {
		return mid
	}

	if direction(mid) {
		return Binary(mid, upper, f, direction)
	} else {
		return Binary(lower, mid, f, direction)
	}
}
