package maumau

import "fmt"

// ExampleNewDeck32 zeigt die Verwendung von NewDeck32.
func ExampleNewDeck32() {
	// Kartenstapel erzeugen
	deck := NewDeck32()

	// Anzahl der Karten ausgeben
	fmt.Println(len(deck.cards))

	// Alle Karten ausgeben
	for _, card := range deck.cards {
		fmt.Println(card)
	}

	// Output:
	// 32
	// ♠7
	// ♠8
	// ♠9
	// ♠10
	// ♠J
	// ♠Q
	// ♠K
	// ♠A
	// ♥7
	// ♥8
	// ♥9
	// ♥10
	// ♥J
	// ♥Q
	// ♥K
	// ♥A
	// ♦7
	// ♦8
	// ♦9
	// ♦10
	// ♦J
	// ♦Q
	// ♦K
	// ♦A
	// ♣7
	// ♣8
	// ♣9
	// ♣10
	// ♣J
	// ♣Q
	// ♣K
	// ♣A

}
