package triangles

import "fmt"

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein gleichschenkliges, rechtwinkliges Dreieck mit diesen Seitenlängen auf der Konsole.
// Der rechte Winkel soll links unten liegen und die Seiten sollen
// vertikal bzw. horizontal verlaufen.
// Die Zeichen für Rand und Füllung des Rechtecks werden als Parameter erwartet.
func DrawTriangle(length int, inner, outer string) {
	for i := 0; i <= length; i++ {
		if i == 0 || i == length {
			for y := 0; y < i; y++ {
				fmt.Print(outer)
			}
			if i == 0 {
				fmt.Println()
			}
		}
		if i != 0 && i != length {

			for y := 0; y < i; y++ {
				if y == 0 || y == i-1 {
					fmt.Print(outer)
				} else {
					fmt.Print(inner)
				}

			}

			fmt.Println()
		}
	}
}
