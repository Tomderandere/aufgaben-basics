package count

// Erwartet eine Liste von Strings sowie einen String, der gezählt werden soll.
// Liefer die Anzahl der Vorkommen des gesuchten Strings in der Liste.
func Count(strings []string, search string) int {
	var counter int
	for i := 0; i < len(strings); i++ {
		if strings[i] == search {
			counter++
		}
	}
	return counter
}
