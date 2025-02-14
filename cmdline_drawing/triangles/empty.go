package triangles

import "fmt"

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein gleichschenkliges, rechtwinkliges Dreieck mit diesen Seitenlängen auf der Konsole.
// Das Dreieck soll komplett mit `#`-Zeichen gefüllt sein.
// Der rechte Winkel soll links unten liegen und die Seiten sollen
// vertikal bzw. horizontal verlaufen.
// Der Rand des Dreiecks soll aus `#`-Zeichen bestehen, der Innenraum soll leer sein.
func DrawEmptyTriangle(length int) {
	for i := 0; i <= length; i++ {
		if i == 0 || i == length {
			for y := 0; y < i; y++ {
				fmt.Print("#")
			}
			if i == 0 {
				fmt.Println()
			}
		}
		if i != 0 && i != length {

			for y := 0; y < i; y++ {
				if y == 0 || y == i-1 {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}

			}

			fmt.Println()
		}
	}
}

// REMARKS
// - Diese Aufgabe ist eine etwas komplexere Variante der Aufgabe "Gefülltes Dreieck" bzw. "Leeres Rechteck".
