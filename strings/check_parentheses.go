package strings

// Erwartet einen String und prüft, ob er korrekte Klammerpaare enthält.
// D.h. der Eingabestring enthält Klammern '(' und ')'.
// Die Funktion soll nun prüfen, ob der String für jede öffnende Klammer auch eine
// schließende Klammer enthält.
// Außerdem darf es keine schließenden Klammern geben, für die es nicht vorher eine
// passende öffnende Klammer gegeben hat.
// Die Funktion soll true liefern, falls der String korrekt geklammert ist.
func CheckParentheses(s string) bool {
	kz := ")"
	ka := "("
	Klammern := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ka[0] {
			Klammern++
		}
		if s[i] == kz[0] {
			Klammern--
		}
	}

	return Klammern == 0
}
