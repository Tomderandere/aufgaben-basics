package make_unique

import "fmt"

// Erwartet eine Liste von Strings.
// Hängt Zahlen an alle mehrfach vorkommenden Strings an, um sie eindeutig zu machen.
func MakeUnique(strings []string) {
	m := make(map[string]int)
	for i := 0; i < len(strings); i++ {
		m[strings[i]]++
		if m[strings[i]] > 1 {
			strings[i] = fmt.Sprintf("%s%s%d", strings[i], "_", m[strings[i]])
		}
	}
}

// REMARKS
// - Dies ist eine ähnliche Aufgabe wie das Zählen von Vorkommen in einer Liste.
//   Die innere Schleife macht i.W. das gleiche wie die Funktion `Count` aus der Aufgabe `count`.
//   Zusätzlich wird der Wert von `count` an den String angehängt, um ihn eindeutig zu machen.
