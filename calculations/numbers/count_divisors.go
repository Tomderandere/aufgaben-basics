package numbers

// Erwartet eine Zahl n >= 1 und liefert die Anzahl der Teiler dieser Zahl zurück.
func CountDivisors(n int) int {
	zähler := 0
	for i := 1; i <= n; i++ {
		t := n % i
		if t == 0 {
			zähler++
		}
	}
	return zähler
}
