package strings

// Erwartet einen String s und liefert s rückwärts zurück.
func Reverse(s string) string {
	var reverse string
	for i := (len(s) - 1); i >= 0; i-- {
		reverse = string(reverse) + string(s[i])
	}
	return reverse
}

// Erwartet zwei Strings s1 und s2 und prüft, ob der eine der andere umgedreht ist.
func IsReverse(s1, s2 string) bool {
	return s2 == Reverse(s1)

}

// Erwartet einen String s und prüft, ob s ein Palindrom ist.
// Ein Palindrom ist ein String, der vorwärts und rückwärts gelesen gleich ist.
func IsPalindrome(s string) bool {

	return s == Reverse(s)

}
