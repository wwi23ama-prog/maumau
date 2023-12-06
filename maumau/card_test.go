package maumau

import "fmt"

// ExampleCardString zeigt die Verwendung von Card.String().
// Die Funktion erzeugt eine Karte und gibt sie als String aus.
func ExampleCard_String() {
	// Karte erzeugen: Pik 7
	c := Card{suit: Spades, rank: Seven}

	// Karte als String ausgeben
	fmt.Println(c.String())

	// Output:
	// ♠7
}
