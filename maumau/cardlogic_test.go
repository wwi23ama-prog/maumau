package maumau

import "fmt"

// ExampleCardsMatch zeigt die Verwendung von CardsMatch.
// Die Funktion erzeugt mehrere Karten und vergleicht
// jeweils zwei davon mittels CardsMatch.
func ExampleCardsMatch() {
	// Karten erzeugen:
	// - Pik 7
	// - Pik 10
	// - Herz 10

	c1 := Card{suit: Spades, rank: Seven}
	c2 := Card{suit: Spades, rank: Ten}
	c3 := Card{suit: Hearts, rank: Ten}

	// Vergleiche Karten:
	// - Pik 7 und Pik 10
	fmt.Println(CardsMatch(c1, c2)) // true

	// - Pik 7 und Herz 10
	fmt.Println(CardsMatch(c1, c3)) // false

	// - Pik 10 und Herz 10
	fmt.Println(CardsMatch(c2, c3)) // true

	// Output:
	// true
	// false
	// true
}
