package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
func MinList(nums []int) int {
	// Basisfall: Wenn die Liste leer ist, gebe 0 zurück
	if len(nums) == 0 {
		return 0
	}

	// Rekursiver Fall: Das Minimum des ersten Elements und des Minimums des restlichen Teils der Liste
	first := nums[0]
	rest := MinListRecursive(nums[1:])

	// Vergleiche das erste Element mit dem Minimum des Restes
	if first < rest || rest == 0 {
		return first
	}
	return rest
}
