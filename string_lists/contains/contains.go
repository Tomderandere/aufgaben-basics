package contains

// Erwartet einen String und eine Liste von Strings.
// Gibt true zurück, wenn der String in der Liste enthalten ist, ansonsten false.
func Contains(strings []string, search string) bool {
	for i := 0; i < len(strings); i++ {
		if strings[i] == search {
			return true
		}
	}
	return false
}
