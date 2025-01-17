package strings

import (
	"slices"
	"strings"
)

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
func IsAnagram(s1, s2 string) bool {
	if len(s1) == len(s2) {
		buchstaben1 := make([]byte, len(s1))
		buchstaben2 := make([]byte, len(s2))
		for i := 0; i < (len(s1)); i++ {
			buchstaben1[i] = s1[i]
			buchstaben2[i] = s2[i]
		}
		slices.Sort(buchstaben1)
		slices.Sort(buchstaben2)
		for i := 0; i < len(buchstaben1); i++ {
			if buchstaben1[i] != buchstaben2[i] {
				return false
			}
		}
		return true
	}
	return false
}

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
// Diese Funktion soll keinen Unterschied zwischen Groß- und Kleinschreibung machen.
func IsAnagramIgnoreCase(s1, s2 string) bool {
	if len(s1) == len(s2) {
		s1 = strings.ToLower(s1)
		s2 = strings.ToLower(s2)
		buchstaben1 := make([]byte, len(s1))
		buchstaben2 := make([]byte, len(s2))
		for i := 0; i < (len(s1)); i++ {
			buchstaben1[i] = s1[i]
			buchstaben2[i] = s2[i]
		}
		slices.Sort(buchstaben1)
		slices.Sort(buchstaben2)
		for i := 0; i < len(buchstaben1); i++ {
			if buchstaben1[i] != buchstaben2[i] {
				return false
			}
		}
		return true
	}
	return false
}
