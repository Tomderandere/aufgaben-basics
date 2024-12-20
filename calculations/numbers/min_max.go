package numbers

// Berechnet das Minimum von zwei Zahlen.
func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// Berechnet das Maximum von zwei Zahlen.
func Max(a, b int) int {
	if a > b {
		return b
	}
	return a
}
