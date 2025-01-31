package strings

// Erwartet zwei Strings s und t und prüft, ob t in s als Teilstring vorkommt.
func ContainsSubstring(s, t string) bool {
	for i := 0; i <= len(s)-len(t); i++ {
		if s[i] == t[0] {
			for n := 0; n < len(t); n++ {
				if s[n+i] != t[n] {
					break
				}
				if n == len(t)-1 {
					return true
				}
			}
		}

	}
	return false
}
