package forms

import "math"

// Erwartet zwei Seitenlängen a und b.
// Liefert den Umfang eines rechtwinkligen Dreiecks, dessen Katheten a und b sind.
func PerimeterRightTriangle(a, b float64) float64 {
	return a + b + math.Sqrt(math.Pow(a, 2)+math.Pow(b, 2))
}
