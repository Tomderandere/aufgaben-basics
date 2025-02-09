package strings

// Erwartet zwei strings s1 und s2.
// Liefert einen String, der abwechselnd aus den Buchstaben des einen und des anderen
// Strings besteht.
func Zip(s1, s2 string) string {
	var lenmax int
	if len(s1) >= len(s2) {
		lenmax = len(s1)
	} else {
		lenmax = len(s2)
	}

	var result string

	for i := 0; i < lenmax; i++ {
		if i < len(s1) {
			result = result + string(s1[i])
		}
		if i < len(s2) {
			result = result + string(s2[i])
		}
	}
	return result
}
