package strings

// Erwartet zwei Strings s und sep sowie eine Zahl n.
// Liefert einen String, der aus n Kopien von s besteht, die duch sep getrennt werden.
func ConcatN(s, sep string, n int) string {
	var result string
	for i := n; i > 0; i-- {
		result = result + s
		for i > 1 {
			result = result + sep
			break
		}
	}
	return result
}
