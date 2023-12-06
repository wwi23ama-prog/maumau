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

// ExampleCardPicture zeigt die Verwendung von Card.Picture().
// Die Funktion erzeugt eine Karte und gibt sie als Unicode-Grafik aus.
func ExampleCard_Picture() {
	// Karte erzeugen: Pik 7
	c := Card{suit: Spades, rank: Seven}

	// Sonderfall: 10 (zwei Zeichen)
	c2 := Card{suit: Spades, rank: Ten}

	// Karten als Unicode-Grafik ausgeben
	fmt.Println(c.Picture())
	fmt.Println(c2.Picture())

	// Output:
	// ┌───────┐
	// │♠7     │
	// │       │
	// │       │
	// │     ♠7│
	// └───────┘
	// ┌───────┐
	// │♠10    │
	// │       │
	// │       │
	// │    ♠10│
	// └───────┘
}
