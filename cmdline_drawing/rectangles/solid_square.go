package rectangles

import "fmt"

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein Quadrat mit dieser Seitenlänge auf der Konsole.
// Das Quadrat soll komplett mit `#`-Zeichen gefüllt sein.
func DrawSolidSquare(length int) {
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
