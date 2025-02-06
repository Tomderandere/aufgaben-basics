package strings

// Erwartet einen String s und liefert einen neuen String,
// in dem jeder Buchstabe aus s zweimal hintereinander vorkommt.
func DuplicateChars(s string) string {
	var result string
	for n := 0; n < len(s); n++ {
		result = string(result) + string(s[n]) + string(s[n])
	}
	return result
}
