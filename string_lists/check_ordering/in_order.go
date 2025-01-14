package check_ordering

// Erwartet eine Liste von Strings und zwei einzelne Strings.
// Überprüft, ob die beiden Strings in der Liste in der gegebenen Reihenfolge vorkommen.
// Gibt `true` zurück, wenn das der Fall ist, ansonsten `false`.
func CheckOrdering(strings []string, first, second string) bool {
	var ifirst int = -10
	var isecond int = -10
	if len(strings) != 0 {
		for i := 0; i < (len(strings)); i++ {
			if strings[i] == first {
				ifirst = i
			}
			if strings[i] == second {
				isecond = i
			}
		}
	}
	if ifirst != -10 && isecond != -10 && ifirst < isecond {
		return true
	}
	return false
}

// REMARKS
// - Diese Aufgabe ist eine komplexere Variante der Aufgabe "Prüfen, ob ein Element in einer Liste vorkommt".
// - Sie können die Lösung der einfachen Variante als Grundlage verwenden und diese entsprechend erweitern.
